 import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './main-container/login/login.component';
import { AdminLoginComponent } from './main-container/admin-login/admin-login.component';
import { AdminHomeComponent } from './main-container/admin-home/admin-home.component';
import { AdminStatsComponent } from './main-container/admin-stats/admin-stats.component';
import { KeeperBookingsComponent } from './main-container/keeper-bookings/keeper-bookings.component';
import { VisitorFindKeepersComponent } from './main-container/visitor-find-keepers/visitor-find-keepers.component';
import { KeeperHomeComponent } from './main-container/keeper-home/keeper-home.component';
import { OwnerHomeComponent } from './main-container/owner-home/owner-home.component';
import { AuthGuard } from './services/auth.guard';
import { OwnerAddPetComponent } from './main-container/owner-add-pet/owner-add-pet.component';
import { OwnerFindKeepersComponent } from './main-container/owner-find-keepers/owner-find-keepers.component';
import { OwnerBookingsComponent } from './main-container/owner-bookings/owner-bookings.component';
import { KeepersReviewComponent } from './main-container/keepers-review/keepers-review.component';
import { OwnerMessagesComponent } from './main-container/owner-messages/owner-messages.component';
import { KeeperMessagesComponent } from './main-container/keeper-messages/keeper-messages.component';


const routes: Routes = [
  // Existing routes...

  // New route
  { path: 'register', component: RegisterComponent },
  { path: 'login', component: LoginComponent },
  { path: 'admin-login', component: AdminLoginComponent },
  { path: 'admin-home', component: AdminHomeComponent },
  { path: 'admin/stats', component: AdminStatsComponent},
  { path: 'keeper/:user_id/bookings', component: KeeperBookingsComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' },
  { path: 'visitor-find-keepers', component: VisitorFindKeepersComponent },
  { path: 'keeper/:user_id', component: KeeperHomeComponent ,canActivate: [AuthGuard]},
  { path: 'owner/:user_id', component: OwnerHomeComponent ,canActivate: [AuthGuard]},
  { path: 'owner/:user_id/add-pet', component: OwnerAddPetComponent ,canActivate: [AuthGuard]},
  { path: 'owner/:user_id/find-keeper', component: OwnerFindKeepersComponent ,canActivate: [AuthGuard]},
  { path: 'owner/:user_id/bookings', component: OwnerBookingsComponent ,canActivate: [AuthGuard]},
  { path: 'keeper/:user_id/reviews', component: KeepersReviewComponent ,canActivate: [AuthGuard]},
  { path: 'owner/:user_id/messages', component: OwnerMessagesComponent ,canActivate: [AuthGuard]},
  { path: 'keeper/:user_id/messages', component: KeeperMessagesComponent ,canActivate: [AuthGuard]},
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
