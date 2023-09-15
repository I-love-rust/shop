import { env } from "$env/dynamic/public";
import type { StdError } from "./error";
import type { RefreshResponse } from "./auth";
import { clearAccessToken, clearRefreshToken, setAccessToken, setRefreshToken } from "$lib/stores/store";



type CallParams = {
    path: string;
    method: "GET" | "POST" | "PUT" | "DELETE";
    body?: object;
}

class RestController {
    private url: string;
    
    constructor(url: string) {
        this.url = url;
    }

    async call<T>(params: CallParams): Promise<T> {
        const res = await fetch(`${this.url}${params.path}`, {
            method: params.method,
            headers: new Headers({
                "Content-Type": "application/json",
            }),
            body: JSON.stringify(params.body)
        })
    
        const json = await res.json()
    
        if (!res.ok) {
            return Promise.reject(json as StdError)
        }
    
        return Promise.resolve(json as T);
    }
    
    async authCall<T>(params: CallParams): Promise<T> {

        const access_token = localStorage.getItem('access_token')
        const refresh_token = localStorage.getItem('refresh_token')

        if (access_token === null || refresh_token === null) {
            return Promise.reject("unauthorized")
        }

        const res = await fetch(`${this.url}${params.path}`, {
            method: params.method,
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": `Bearer ${access_token}`
            }),
            body: JSON.stringify(params.body)
        })
    
        const json = await res.json()
    
        // refresh access token
        if (res.status === 401) {
            const req = await fetch(`${this.url}/auth/refresh`, {
                method: "POST",
                headers: new Headers({
                    "Content-Type": "application/json",
                }),
                body: JSON.stringify({
                    refresh_token: refresh_token,
                })
            })
            const json = await req.json()

            if (req.status === 401) {
                clearRefreshToken()
                clearAccessToken()
                return Promise.reject(json as Error)
            }

            const res = await Promise.resolve(json as RefreshResponse)
            setAccessToken(res.access_token)
            setRefreshToken(res.refresh_token)

            return this.authCall(params)
        }
    
        if (!res.ok) {
            return Promise.reject(json as StdError)
        }
    
        return Promise.resolve(json as T);
    }
}

export const restController = new RestController(env.PUBLIC_RESTAPI_URL);