import { Component } from '@angular/core';
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

  constructor(private authService: AuthService) {

  }

  onSignUp() {
    console.log(this.username, this.password, this.email)
    this.authService.register(this.username, this.password, this.email).subscribe(data => console.log("success", data), err => console.log("err", err))
  }
}
