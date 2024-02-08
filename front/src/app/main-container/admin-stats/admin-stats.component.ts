import { Component, OnInit } from '@angular/core';
import { AdminService } from 'src/app/services/admin.service';
import { ChartModule } from 'primeng/chart';
import { AuthService } from 'src/app/services/auth.service';


interface petCount{
  cats: number;
  dogs: number;
}

interface userCount{
  owners: number;
  keepers: number;
}

interface money{
  keeper: number;
  app: number;
}

@Component({
  selector: 'app-admin-stats',
  templateUrl: './admin-stats.component.html',
  styleUrl: './admin-stats.component.css'
})
export class AdminStatsComponent implements OnInit{

  constructor(private adminService: AdminService, private authService : AuthService) {}

  petCount: petCount = {cats: 0, dogs: 0};
  userCount: userCount = {owners: 0, keepers: 0};
  money: money = {keeper: 0, app: 0};

  petsData: any;
  petOptions: any;

  usersData: any;
  usersOptions: any;

  moneyData: any;
  moneyOptions: any;

  ngOnInit() {
    this.adminService.getPetCount().subscribe((data : petCount) => {
      this.petCount = data;
      this.updateChartPetsData();
    });
    this.adminService.getUserCount().subscribe((data : userCount) => {
      this.userCount = data;
      this.updateChartUsersData();
    });
    this.adminService.getMoney().subscribe((data : money) => {
      this.money = data;
      this.updateChartMoneyData();
    });

  }

  updateChartPetsData() {
    const documentStyle = getComputedStyle(document.documentElement);
    const textColor = documentStyle.getPropertyValue('--text-color');


    this.petsData = {
      labels: ['Cats', 'Dogs'],
      datasets: [
          {
              data: [this.petCount.cats, this.petCount.dogs],
              backgroundColor: [documentStyle.getPropertyValue('--blue-500'), documentStyle.getPropertyValue('--yellow-500'), documentStyle.getPropertyValue('--green-500')],
              hoverBackgroundColor: [documentStyle.getPropertyValue('--blue-400'), documentStyle.getPropertyValue('--yellow-400'), documentStyle.getPropertyValue('--green-400')]
          }
      ]
    };


    this.petOptions = {
        cutout: '60%',
        plugins: {
            legend: {
                labels: {
                    color: textColor
                }
            }
        }
    };
  }

  updateChartUsersData() {
        const documentStyle = getComputedStyle(document.documentElement);
        const textColor = documentStyle.getPropertyValue('--text-color');
        const textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
        const surfaceBorder = documentStyle.getPropertyValue('--surface-border');
        
        console.log(this.userCount.owners);

        this.usersData = {
            labels: ['Users'],
            datasets: [
                {
                    label: 'Keepers',
                    backgroundColor: documentStyle.getPropertyValue('--blue-500'),
                    borderColor: documentStyle.getPropertyValue('--blue-500'),
                    data: [this.userCount.keepers]
                },
                {
                    label: 'Owners',
                    backgroundColor: documentStyle.getPropertyValue('--pink-500'),
                    borderColor: documentStyle.getPropertyValue('--pink-500'),
                    data: [this.userCount.owners]
                }
            ]
        };

        this.usersOptions = {
            maintainAspectRatio: false,
            aspectRatio: 0.8,
            plugins: {
                legend: {
                    labels: {
                        color: textColor
                    }
                }
            },
            scales: {
                x: {
                    ticks: {
                        color: textColorSecondary,
                        font: {
                            weight: 500
                        }
                    },
                    grid: {
                        color: surfaceBorder,
                        drawBorder: false
                    }
                },
                y: {
                    ticks: {
                        color: textColorSecondary
                    },
                    grid: {
                        color: surfaceBorder,
                        drawBorder: false
                    }
                }

            }
        };
    }

    updateChartMoneyData() {
      const documentStyle = getComputedStyle(document.documentElement);
        const textColor = documentStyle.getPropertyValue('--text-color');
        this.moneyData = {
            labels: ['keepers','app'],
            datasets: [
                {
                    data: [this.money.keeper, this.money.app],
                    backgroundColor: [documentStyle.getPropertyValue('--blue-500'), documentStyle.getPropertyValue('--yellow-500'), documentStyle.getPropertyValue('--green-500')],
                    hoverBackgroundColor: [documentStyle.getPropertyValue('--blue-400'), documentStyle.getPropertyValue('--yellow-400'), documentStyle.getPropertyValue('--green-400')]
                }
            ]
        };

        this.moneyOptions = {
            plugins: {
                legend: {
                    labels: {
                        usePointStyle: true,
                        color: textColor
                    }
                }
            }
        };
    }
    onLogout() {
        this.authService.logout();
    }
}

  
    
