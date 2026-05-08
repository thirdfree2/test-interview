import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { UserService } from '../../core/services/user.service';
import { ToastService } from '../../core/services/toast.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {
  registerForm: any = null
  private toast = inject(ToastService);
  private userService = inject(UserService)
  private router = inject(Router);
  constructor(
    private fb: FormBuilder
  ) {
    this.registerForm = this.fb.group({
      name: ['', Validators.required],
      password: ['', Validators.required],
      confirmPassword: ['', Validators.required]
    });
  }

  postRegister() {
    if (this.registerForm.invalid) {
      this.registerForm.markAllAsTouched();
      return;
    }

    if (
      this.registerForm.value.password !==
      this.registerForm.value.confirmPassword
    ) {
      this.toast.error('password ไม่ตรงกัน');
      return;
    }

    this.userService.register({
      name: this.registerForm.value.name!,
      password: this.registerForm.value.password!
    }).subscribe({
      next: (value) => {
        this.toast.success(value.message ?? "OK");
        this.router.navigate([
          '/'
        ]);
      },
      error: (err) => {
        this.toast.error(err.error.message ?? "Error");
        if (err.error.data.name) {
          this.toast.error(err.error.data.name ?? "Error");
        }
        if (err.error.data.password) {
          this.toast.error(err.error.data.password ?? "Error");
        }
      },
    })
  }
}
