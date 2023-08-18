import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, catchError, tap } from 'rxjs';
import { environment } from 'src/environments/environment';
import { User } from '../models/user.model';


interface UserResponse {
  user: User;
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  public isAuthenticated: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);

  public user: User | null = null;

  private url = `${environment.apiUrl}/users`;


  constructor(private http: HttpClient) {
    this.isAuthenticated.next(this.loadUser());
  }

  register(username: string, password: string, email: string) {
    return this.http.post<UserResponse>(this.url, { user: { username, password, email } }).pipe(
      tap(data => {
        this.setUser(data.user);
        this.isAuthenticated.next(true);
      }),
    );
  }

  login(password: string, email: string) {
    return this.http.post<UserResponse>(`${this.url}/login`, { user: { password, email } }).pipe(
      tap(data => {
        this.setUser(data.user);
        this.isAuthenticated.next(true);
      }),
    );
  }

  private setUser(user: User) {
    this.user = user;
    localStorage.setItem("user", JSON.stringify(this.user));
  }

  private loadUser(): boolean {
    const rawUser = localStorage.getItem("user");
    if (rawUser === null) {
      return false;;
    }

    this.user = JSON.parse(rawUser);
    return true;
  }
}
