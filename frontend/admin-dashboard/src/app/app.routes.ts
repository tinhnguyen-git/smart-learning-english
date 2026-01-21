import { Routes } from '@angular/router';
import { DashboardLayoutComponent } from './layout/dashboard-layout/dashboard-layout.component';
import { StoryListComponent } from './features/story-approval/story-list/story-list.component';
import { AffiliateManagerComponent } from './features/affiliate/affiliate-manager/affiliate-manager.component';
import { ScraperComponent } from './features/link-scraper/scraper/scraper.component';

export const routes: Routes = [
    { path: 'login', loadComponent: () => import('./features/auth/login/login.component').then(m => m.LoginComponent) },
    { path: 'register', loadComponent: () => import('./features/auth/register/register.component').then(m => m.RegisterComponent) },
    { path: 'profile', loadComponent: () => import('./features/user/profile/profile.component').then(m => m.ProfileComponent) },
    {
        path: '',
        redirectTo: 'dashboard',
        pathMatch: 'full'
    },
    {
        path: 'dashboard',
        component: DashboardLayoutComponent,
        children: [
            { path: '', redirectTo: 'stories', pathMatch: 'full' },
            { path: 'stories', component: StoryListComponent },
            { path: 'affiliates', component: AffiliateManagerComponent },
            { path: 'scraper', component: ScraperComponent }
        ]
    }
];
