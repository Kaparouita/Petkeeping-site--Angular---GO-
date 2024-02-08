import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KeeperMessagesComponent } from './keeper-messages.component';

describe('KeeperMessagesComponent', () => {
  let component: KeeperMessagesComponent;
  let fixture: ComponentFixture<KeeperMessagesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [KeeperMessagesComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(KeeperMessagesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
