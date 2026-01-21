import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthService } from '../../../core/services/auth.service';
import { UserService } from '../../../core/services/user.service';

@Component({
    selector: 'app-profile',
    standalone: true,
    imports: [CommonModule],
    template: `
    <div class="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl">
        <div class="p-8">
          <div class="uppercase tracking-wide text-sm text-indigo-500 font-semibold">User Profile</div>
          <div class="mt-4">
            <h1 class="text-2xl font-bold text-gray-900">{{ user?.full_name || 'Guest' }}</h1>
            <p class="mt-2 text-gray-600">{{ user?.email }}</p>
            <p class="mt-1 text-sm text-gray-500">User ID: {{ user?.id }}</p>
          </div>

          <div class="mt-8 border-t pt-6">
            <div *ngIf="message" class="mb-4 p-4 rounded" [ngClass]="isError ? 'bg-red-100 text-red-700' : 'bg-green-100 text-green-700'">
              {{ message }}
            </div>

            <button (click)="upgradeToPremium()" [disabled]="loading"
              class="w-full bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 disabled:opacity-50">
              <span *ngIf="loading">Processing...</span>
              <span *ngIf="!loading">Upgrade to Premium</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  `
})
export class ProfileComponent implements OnInit {
    user: any = null;
    loading: boolean = false;
    message: string = '';
    isError: boolean = false;

    constructor(private authService: AuthService, private userService: UserService) { }

    ngOnInit() {
        this.user = this.authService.getUser();
    }

    upgradeToPremium() {
        this.loading = true;
        this.message = '';
        this.isError = false;

        this.userService.upgrade().subscribe({
            next: (res) => {
                this.message = 'Successfully upgraded to Premium!';
                this.loading = false;
                // Optionally update local user data if backend returns new status
            },
            error: (err) => {
                this.message = 'Upgrade failed. Please try again.';
                this.isError = true;
                this.loading = false;
                console.error(err);
            }
        });
    }
}
