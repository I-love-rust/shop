import type { UserDto } from "$lib/dto";
import { restController } from "./controller";

export type SignUpParams = {
  username: string;
  email: string;
  password: string;
}

export type SignInParams = {
  login: string;
  password: string;
}


export type SignInResponse = {
  access_token: string;
  refresh_token: string;
}

export type RefreshParams = {
  refresh_token: string;
}

export type RefreshResponse = {
  access_token: string;
  refresh_token: string;
}

export type CheckResponse = {
  status: boolean;
}

// Function to make a GET request to /check
export async function checkStatus(): Promise<CheckResponse> {
  return restController.authCall<CheckResponse>({
    path: "/user/check",
    method: "GET"
  })
}

// Function to make a POST request to /refresh
export async function refresh(params: RefreshParams) {
  return restController.call<RefreshResponse>({
    path: "/auth/refresh",
    method: "POST",
    body: params
  })
}

// Function to make a POST request to /signup
export async function signUp(params: SignUpParams): Promise<UserDto> {
  return restController.call<UserDto>({
    path: "/auth/signup",
    method: "POST",
    body: params
  })
}

// Function to make a POST request to /signin
export async function signIn(params: SignInParams): Promise<SignInResponse> {
  return restController.call<SignInResponse>({
    path: "/auth/signin",
    method: "POST",
    body: params
  })
}