import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KeeperBookingsComponent } from './keeper-bookings.component';

describe('KeeperBookingsComponent', () => {
  let component: KeeperBookingsComponent;
  let fixture: ComponentFixture<KeeperBookingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [KeeperBookingsComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(KeeperBookingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
