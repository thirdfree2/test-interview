import { Injectable, signal } from '@angular/core';

export type ToastType =
    | 'success'
    | 'error'
    | 'warning'
    | 'info';

export interface ToastData {
    message: string;
    type: ToastType;
}

@Injectable({
    providedIn: 'root',
})
export class ToastService {

    toast = signal<ToastData | null>(null);

    show(
        message: string,
        type: ToastType = 'info'
    ) {

        this.toast.set({
            message,
            type,
        });

        setTimeout(() => {
            this.toast.set(null);
        }, 3000);
    }

    success(message: string) {
        this.show(message, 'success');
    }

    error(message: string) {
        this.show(message, 'error');
    }

    warning(message: string) {
        this.show(message, 'warning');
    }

    info(message: string) {
        this.show(message, 'info');
    }
}