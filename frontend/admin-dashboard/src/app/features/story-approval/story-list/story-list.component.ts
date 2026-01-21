import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

interface Story {
  id: string;
  title: string;
  author: string;
  status: 'PENDING' | 'APPROVED' | 'REJECTED';
  submittedAt: Date;
}

@Component({
  selector: 'app-story-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './story-list.component.html',
  styleUrl: './story-list.component.css'
})
export class StoryListComponent {
  stories: Story[] = [
    { id: '1', title: 'The Adventures of Learning English', author: 'John Doe', status: 'PENDING', submittedAt: new Date() },
    { id: '2', title: 'Advanced Grammar Tips', author: 'Jane Smith', status: 'PENDING', submittedAt: new Date(Date.now() - 86400000) },
    { id: '3', title: 'Vocabulary Expansion', author: 'Bob Johnson', status: 'APPROVED', submittedAt: new Date(Date.now() - 172800000) },
  ];

  approve(id: string) {
    const story = this.stories.find(s => s.id === id);
    if (story) story.status = 'APPROVED';
  }

  reject(id: string) {
    const story = this.stories.find(s => s.id === id);
    if (story) story.status = 'REJECTED';
  }
}
