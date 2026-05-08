import { Component, inject } from '@angular/core';
import { ApiService } from '../../core/services/api.service';
import { CommonModule } from '@angular/common';
import { UserService } from '../../core/services/user.service';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ToastService } from '../../core/services/toast.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginForm: any = null
  constructor(
    private fb: FormBuilder
  ) {
    this.loginForm = this.fb.group({
      name: ['', Validators.required],
      password: ['', Validators.required],
    });
  }

  private userService = inject(UserService)
  private toast = inject(ToastService);
  private router = inject(Router);


  login() {
    if (this.loginForm.invalid) {
      this.loginForm.markAllAsTouched();
      return;
    }

    this.userService
      .login({
        name: this.loginForm.value.name!,
        password: this.loginForm.value.password!,
      })
      .subscribe({
        next: (res) => {
          const token = res.data;

          if (token) {
            localStorage.setItem(
              'access_token',
              token
            );
          }

          this.toast.success(
            res.message ?? 'เข้าสู่ระบบสำเร็จ'
          );
          this.router.navigate([
            '/my-profile'
          ]);


        },
        error: (err) => {
          this.toast.error(err.error.message ?? "Error");
        },
      });
  }
}
