import { Component } from '@angular/core';
import { TableModule } from 'primeng/table';
import { UserServiceService } from '../../services/user-service.service';
import { User } from '../../models/user.model';

@Component({
  selector: 'app-visitor-find-keepers',
  templateUrl: './visitor-find-keepers.component.html',
  styleUrl: './visitor-find-keepers.component.css'
})
export class VisitorFindKeepersComponent {
  constructor(private userService : UserServiceService) {}
  keepers: User[] = []; 

  ngOnInit(): void {
    this.userService.getKeepers().subscribe(
      (data) => {
        console.log(data);
        this.keepers = data;
        
      },
      (error) => {
        console.log(error);
      }
    );

    }
  }
