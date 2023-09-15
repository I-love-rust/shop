export type UserDto = {
    id: string;
    username: string;
    name: string;
    avatar: string;
    bio: string;
    email: string;
    role: number;
    created_at: string;
}

export type UserPublicDto = {
    id: number;
    username: string;
    name: string;
    avatar: string;
    bio: string;
    created_at: string;
}