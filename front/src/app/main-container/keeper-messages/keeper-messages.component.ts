import { Component, OnInit } from '@angular/core';
import { UserServiceService } from 'src/app/services/user-service.service';
import { Message } from 'src/app/models/message.model';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/models/user.model';

@Component({
  selector: 'app-keeper-messages',
  templateUrl: './keeper-messages.component.html',
  styleUrl: './keeper-messages.component.css'
})
export class KeeperMessagesComponent {
  keeper: User | undefined;
  keeper_id: number | undefined;
  messages: Message[] = [];

  constructor(private userService: UserServiceService,private route : ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const userId = params.get('user_id') ;
      this.keeper_id = parseInt(userId || '0', 10);
      this.userService.getKeeper(this.keeper_id).subscribe(
        data => {
          this.keeper = data; 
          this.userService.getMessagesByName(this.keeper.username).subscribe(
            data => {
              this.messages = data; 
              console.log(data);
            },
            error => {
              console.error('Error fetching keeper data', error);
            }
          );
        },
        error => {
          console.error('Error fetching keeper data', error);
        }
      );
    });
  }
}
