import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError } from 'rxjs/operators';
import { environment } from 'src/environments/environment';
import { Tags } from '../models/tags.model';

@Injectable({
  providedIn: 'root'
})
export class TagService {
  private url = `${environment.apiUrl}/tags`;

  constructor(private http: HttpClient) { }

  getTags() {
    return this.http.get<Tags>(this.url);
  }
}
