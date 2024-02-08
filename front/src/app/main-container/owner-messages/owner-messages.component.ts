import { Component, OnInit } from '@angular/core';
import { UserServiceService } from 'src/app/services/user-service.service';
import { Message } from 'src/app/models/message.model';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/models/user.model';


@Component({
  selector: 'app-owner-messages',
  templateUrl: './owner-messages.component.html',
  styleUrl: './owner-messages.component.css'
})
export class OwnerMessagesComponent implements OnInit {
    owner: User | undefined;
    owner_id: number | undefined;
    messages: Message[] = [];

    constructor(private userService: UserServiceService,private route : ActivatedRoute) { }

    ngOnInit(): void {
      this.route.paramMap.subscribe(params => {
        const userId = params.get('user_id') ;
        this.owner_id = parseInt(userId || '0', 10);
        this.userService.getOwner(this.owner_id).subscribe(
          data => {
            this.owner = data; 
            this.userService.getMessagesByName(this.owner.username).subscribe(
              data => {
                this.messages = data; 
                console.log(data);
              },
              error => {
                console.error('Error fetching owner data', error);
              }
            );
          },
          error => {
            console.error('Error fetching owner data', error);
          }
        );
      });
    }
}

