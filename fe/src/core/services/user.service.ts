import { inject, Injectable } from "@angular/core";
import { BaseResponse } from "../model/base_reponse";
import { Observable } from "rxjs";
import { ApiService } from "./api.service";
import { LoginRequest, RegisterRequest, UserProfile } from "../model/user_model";


@Injectable({
    providedIn: 'root',
})
export class UserService {
    private api = inject(ApiService);

    login(
        body: LoginRequest
    ): Observable<BaseResponse<string>> {

        return this.api.post<
            BaseResponse<string>
        >('/users/login', body);
    }

    register(body: RegisterRequest): Observable<BaseResponse<any>> {
        return this.api.post<
            BaseResponse<any>
        >('/users/register', body);
    }

    getUserProfile(): Observable<BaseResponse<UserProfile>> {
        return this.api.get<
            BaseResponse<UserProfile>
        >('/profile/me');
    }
}