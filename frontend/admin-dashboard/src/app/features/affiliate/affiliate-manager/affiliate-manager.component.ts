import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

interface AffiliateAd {
  id: string;
  code: string;
  active: boolean;
  impressions: number;
}

@Component({
  selector: 'app-affiliate-manager',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './affiliate-manager.component.html',
  styleUrl: './affiliate-manager.component.css'
})
export class AffiliateManagerComponent {
  newAdCode: string = '';
  activeAds: AffiliateAd[] = [
    { id: '1', code: '<script>...</script>', active: true, impressions: 1250 },
    { id: '2', code: '<iframe...></iframe>', active: false, impressions: 340 }
  ];

  addAd() {
    if (!this.newAdCode.trim()) return;

    this.activeAds.unshift({
      id: Math.random().toString(36).substr(2, 9),
      code: this.newAdCode,
      active: true,
      impressions: 0
    });
    this.newAdCode = '';
  }

  toggleStatus(ad: AffiliateAd) {
    ad.active = !ad.active;
  }
}
