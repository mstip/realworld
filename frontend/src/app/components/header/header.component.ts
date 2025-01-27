import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  public isAuthenticated = false;
  public username = "";

  constructor(private authService: AuthService) {
  }

  ngOnInit(): void {
    this.authService.isAuthenticated.subscribe(isAuthenticated => {
      this.isAuthenticated = isAuthenticated;
      if(!this.isAuthenticated) {
        this.username = "";
        return;
      }
      this.username = this.authService.user?.username === undefined ? "" : this.authService.user?.username;
    }
    )
  }
}
