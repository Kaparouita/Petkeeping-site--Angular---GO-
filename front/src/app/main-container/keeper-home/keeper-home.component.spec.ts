import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KeeperHomeComponent } from './keeper-home.component';

describe('KeeperHomeComponent', () => {
  let component: KeeperHomeComponent;
  let fixture: ComponentFixture<KeeperHomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [KeeperHomeComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(KeeperHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
