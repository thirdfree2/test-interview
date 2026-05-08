export interface LoginRequest {
    name: string;
    password: string;
}

export interface RegisterRequest {
    name: string;
    password: string;
}

export interface UserProfile {
    id: number,
    name: string,
    created_at: string,
    updated_at: string
}