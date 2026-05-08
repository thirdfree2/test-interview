import { Component, inject } from '@angular/core';
import { UserService } from '../../core/services/user.service';
import { UserProfile } from '../../core/model/user_model';
import { AuthService } from '../../core/services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-my-profile',
  imports: [],
  templateUrl: './my-profile.component.html',
  styleUrl: './my-profile.component.css'
})
export class MyProfileComponent {
  private userService = inject(UserService)
  private authService = inject(AuthService)
  private router = inject(Router);
  userProfile?: UserProfile = {
    id: 0,
    name: '',
    created_at: '',
    updated_at: ''
  }
  constructor() {
    this.getUserProfile()
  }

  getUserProfile() {
    this.userService.getUserProfile().subscribe({
      next: (res) => {
        console.log(res);
        this.userProfile = res.data
      },
      error(err) {
        console.log(err)
      },
    })
  }

  clearToken() {
    this.authService.clearToken()
    this.router.navigate([
      '/'
    ]);
  }
}
