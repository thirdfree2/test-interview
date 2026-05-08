import { Routes } from '@angular/router';
import { LoginComponent } from '../pages/login/login.component';
import { MyProfileComponent } from '../pages/my-profile/my-profile.component';
import { RegisterComponent } from '../pages/register/register.component';

export const routes: Routes = [
    {
        path: '',
        component: LoginComponent,
    },
    {
        path: 'register',
        component: RegisterComponent,
    },
    {
        path: 'my-profile',
        component: MyProfileComponent,
    },
];
