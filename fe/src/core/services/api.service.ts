import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { API_URL } from '../constant/api';


@Injectable({
    providedIn: 'root',
})
export class ApiService {
    private http = inject(HttpClient);

    get<T>(endpoint: string) {
        return this.http.get<T>(`${API_URL}${endpoint}`);
    }

    post<T>(endpoint: string, body: any) {
        return this.http.post<T>(`${API_URL}${endpoint}`, body);
    }

    put<T>(endpoint: string, body: any) {
        return this.http.put<T>(`${API_URL}${endpoint}`, body);
    }

    delete<T>(endpoint: string) {
        return this.http.delete<T>(`${API_URL}${endpoint}`);
    }
}