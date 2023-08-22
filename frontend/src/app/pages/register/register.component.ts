import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  public username: string = "";
  public email: string = "";
  public password: string = "";
  public error: string | null = null;

  constructor(private authService: AuthService, private router: Router) {
  }

  onSignUp() {
    this.error = null;
    this.authService.register(this.username, this.password, this.email).subscribe(data => {
      this.router.navigate(['/']);
    }, err => {
      this.error = JSON.stringify(err.error.errors);
    })
  }
}
