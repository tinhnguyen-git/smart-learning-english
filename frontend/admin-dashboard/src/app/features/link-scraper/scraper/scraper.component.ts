import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { StoryService } from '../../../core/services/story.service';

@Component({
  selector: 'app-scraper',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './scraper.component.html',
  styleUrl: './scraper.component.css'
})
export class ScraperComponent {
  url: string = '';
  loading: boolean = false;
  result: any = null;
  errorMessage: string = '';

  constructor(private storyService: StoryService) { }

  scrape() {
    if (!this.url) return;

    this.loading = true;
    this.result = null;
    this.errorMessage = '';

    this.storyService.scrape(this.url).subscribe({
      next: (data) => {
        this.result = data;
        this.loading = false;
      },
      error: (err) => {
        this.loading = false;
        this.errorMessage = 'Failed to scrape URL. Please try again.';
        console.error(err);
      }
    });
  }

  importToStory() {
    alert('Imported to pending stories!');
    this.result = null;
    this.url = '';
  }
}
