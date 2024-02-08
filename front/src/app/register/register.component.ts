import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserServiceService } from '../services/user-service.service';
import { User } from '../models/user.model';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrl: './register.component.css',
})
export class RegisterComponent {
  constructor(
    private userService: UserServiceService,
    private router: Router,
    private datePipe: DatePipe
  ) {}

  username: string = '';
  email: string = '';
  password: string = '';
  confirmPassword: string = '';
  firstname: string = '';
  lastname: string = '';
  birthdate: Date = new Date();
  spaceType: string = '';
  gender: string = '';
  type: string = ''; // keeper or owner
  country: string = 'Greece';
  city: string = '';
  address: string = '';
  personalpage: string = '';
  job: string = '';
  phone: string = '';
  showPassword: boolean = false;
  showConfirmPassword: boolean = false;
  passwordValidationMessage: string = '';
  passwordInvalid: boolean = false;
  confirmPasswordValidationMessage: string = '';

  hasCat: boolean = false;
  hasDog: boolean = false;
  catPrice: number = 8;
  dogPrice: number = 8;
  spaceDescription: string = '';

  showMap: boolean = false;
  lat: number = 1006;
  lon: number = 1966; // default values

  //----------------------VALIDATION----------------------//

  onCountrySelected(country: string) {
    this.country = country;
  }

  // !TODO: Change this to the lat/lon coordinates
  onMap(address: string) {
    this.address = address;
  }

  dogCheckboxChange() {
    // Do something
  }

  catCheckboxChange() {
    // Do something
  }


  checkPassword() {
    const { message, valid } = this.validatePassword(
      this.password,
      this.confirmPassword
    );
    this.passwordValidationMessage = message;
    this.passwordInvalid = valid === 'danger';
    return valid;
  }

  toggleConfirmPasswordVisibility() {
    this.showConfirmPassword = !this.showConfirmPassword;
  }

  togglePasswordVisibility() {
    this.showPassword = !this.showPassword;
  }

  contains50PercentNumbers(password: string) {
    const numbers = password.match(/\d/g);
    if (!numbers) {
      return false;
    }
    return numbers.length / password.length >= 0.5;
  }

  containsKeyWord(password: string) {
    const keyWords = ['dog', 'cat', 'gata', 'skylos'];
    for (var i = 0; i < keyWords.length; i++) {
      if (password.includes(keyWords[i])) {
        return true;
      }
    }
    return false;
  }

  validateForm() {
    let resp = this.validatePassword(this.password, this.confirmPassword);
    if (resp.valid === 'danger') {
      return false;
    }
    return true;
  }

  containsCapLetterSymbolNumber(password: string) {
    const regex = /^(?=.*[A-Z])(?=.*[!@#$&*])(?=.*\d).+$/;
    return regex.test(password);
  }

  validatePassword(password: string, confirmPassword: string) {
    var pattern = /^(?=.*[!@#$%^&*])(?=.*[a-zA-Z])(?=.*\d).*$/;

    if (password !== confirmPassword) {
      return { message: 'Passwords do not match', valid: 'danger' };
    } else if (password.length < 8) {
      return { message: 'Password is too short', valid: 'danger' };
    } else if (password.length > 15) {
      return { message: 'Password is too long', valid: 'danger' };
    } else if (!pattern.test(password)) {
      return {
        message:
          'Password contains invalid characters(at least one special symbol,letter,number)',
        valid: 'danger',
      };
    } else if (this.contains50PercentNumbers(password)) {
      return { message: 'Password contains 50% numbers', valid: 'danger' };
    } else if (this.containsKeyWord(password)) {
      return { message: 'Password contains a keyword', valid: 'danger' };
    } else if (this.containsCapLetterSymbolNumber(password)) {
      return { message: 'Password is strong', valid: 'success' };
    }
    return { message: 'Password is medium', valid: 'warning' };
  }

  formatDate(date: Date): string {
    let formatBirthday = this.datePipe.transform(
      date,
      'yyyy-MM-ddTHH:mm:ssZZZZZ'
    );

    return formatBirthday as string;
  }

  checkSpaceType(space: string) {
    if (space === 'outdoor') {
      return false;
    }
    return true;
  }

  ngOnInit(): void {}

  saveUser() {
    //format birthday
    let formatBirthday = this.formatDate(this.birthdate);
    console.log(this.country);

    const user = new User(this.username,this.email,this.password,this.firstname,this.lastname,formatBirthday,this.gender,this.type,this.country,
    this.city,this.address,this.personalpage,this.job,this.phone,this.lat,this.lon);


    if (this.type === 'keeper') {
      user.space_type = this.spaceType;
      user.cat_keep = this.hasCat;
      user.dog_keep = this.hasDog;
      user.cat_price = this.catPrice;
      user.dog_price = this.dogPrice;
      user.space_desc = this.spaceDescription;

      this.userService.saveKeeper(user).subscribe(
        (data) => {
          console.log(data);
        },
        (error) => {
          console.log(error);
          return false;
        }
      );
    } else if (this.type === 'owner') {
      this.userService.saveOwner(user).subscribe(
        (data) => {
          console.log(data);
        },
        (error) => {
          console.log(error);
          return false;
        }
      );
    }
    return true
  }

  onSubmit() {
    console.log('register form submitted');
    if (this.validateForm()) {
      if (!this.saveUser()){
        console.log('invalid form');
        return;
      }

      console.log('valid form');
      this.router.navigate(['/login']);
    } else {
      console.log('invalid form');
    }
  }
}
