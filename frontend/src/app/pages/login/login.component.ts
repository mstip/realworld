import { Component } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  public email: string = "";
  public password: string = "";

  constructor(private authService: AuthService) {

  }

  onSignIn() {
    this.authService.login(this.password, this.email).subscribe(data => console.log("success", data), err => console.log("err", err))
  }
}
