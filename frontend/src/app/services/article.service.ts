import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Articles } from '../models/articles.model';

@Injectable({
  providedIn: 'root'
})
export class ArticleService {
  private url = `${environment.apiUrl}/articles`;

  
  constructor(private http: HttpClient) { }

  getArticles() {
    return this.http.get<Articles>(this.url);
  }
}
