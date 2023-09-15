import type { UserPublicDto } from "$lib/dto";
import { restController } from "./controller";

export type UploadParams = {
    image: string;
}
  
export type UploadResponse = {
    path: string;
}

export type getUserParams = {
    username: string;
}

// Function to make a GET request to /upload
export async function getUserByUsername(params: getUserParams): Promise<UserPublicDto> {
    return restController.authCall<UserPublicDto>({
        path: `/user/by_username?username=${params.username}`,
        method: "GET",
    })
}

export async function uploadImage(params: UploadParams): Promise<UploadResponse> {
    return restController.authCall<UploadResponse>({
        path: "/user/upload",
        method: "POST",
        body: params
    })
}