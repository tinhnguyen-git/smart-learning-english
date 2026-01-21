import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class StoryService {
    private apiUrl = 'http://localhost:8080/stories';

    constructor(private http: HttpClient) { }

    scrape(url: string): Observable<any> {
        return this.http.post(`${this.apiUrl}/scrape`, { url });
    }
}
