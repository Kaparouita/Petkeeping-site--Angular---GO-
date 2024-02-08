import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserServiceService} from '../../services/user-service.service';
import { AuthService } from 'src/app/services/auth.service';
import { LoginResp } from 'src/app/models/login.model';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  username: string = '';
  password: string = '';

  constructor(private userService : UserServiceService,private router : Router,private authService : AuthService) {}

  onLogin() {
    let params = new LoginResp(this.username, this.password, '', 0);

    this.authService.login(params).subscribe(
      (data) => {
        if(data.user_type == 'keeper'){
        console.log(data);
        let url = '/keeper/' + data.user_id ;
        console.log(url);
        this.router.navigate([url]);
        }else if(data.user_type == 'owner'){
          console.log(data);
          let url = '/owner/' + data.user_id ;
          console.log(url);
          this.router.navigate([url]);
        }
      },
      (error) => {
        console.log(error);
      }
    );
  }
}
