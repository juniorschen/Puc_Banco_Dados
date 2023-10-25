import { Component, OnInit } from '@angular/core'

import { Item } from './item'
import { ItemService } from './item.service'
import { RouterExtensions } from '@nativescript/angular';
import { SignalrCore } from 'nativescript-signalr-core';

@Component({
  selector: 'ns-items',
  templateUrl: './items.component.html',
})
export class ItemsComponent implements OnInit {

  signalrCore = new SignalrCore();
  items: Array<Item> = new Array<Item>();

  constructor(private itemService: ItemService, private routerExtension: RouterExtensions) {
    console.log('cccd')
    /* this.signalrCore.on('notify', (data) => {
      console.log('notificacao recebida', data);
    });

    this.signalrCore.on('onReceivePendingMessages', (data) => {
      console.log('onReceivePendingMessages')
      var ids = new Array<string>();
      data.arguments.forEach(element => {
        element.notifications.forEach(noti => {
          ids.push(noti.notificationId);
        });
      });
      this.consumeMessages(ids);
    });

    this.signalrCore.start('http://amos-junior.com:9999/notifications/notification').then((isConnected: boolean) => {
      console.log('isConnected? ', isConnected);
    }); */
  }

  async ngOnInit(): Promise<void> {
    this.items = await this.itemService.getItems();
  }

  addRobozinho() {
    this.routerExtension.navigate(["/item"]);
  }

  invoke() {
    this.signalrCore.invoke('AddToGroupWithTenantAndEnvironmentAndAppId', ...['C729D2B2-F0CD-4E72-A4BC-E06B5603B60E', 'EE19958D-5F6A-4201-B989-08D7468F1643', 'd60577e2-b84f-40a0-0829-08db4c068864', 'MOB02']);
    this.signalrCore.invoke('GetPendingMessagesWithTenantAndEnvironmentAndAppId', ...['C729D2B2-F0CD-4E72-A4BC-E06B5603B60E', 'EE19958D-5F6A-4201-B989-08D7468F1643', 'd60577e2-b84f-40a0-0829-08db4c068864', 'MOB02']);
  }

  consumeMessages(ids) {
    console.log('consumeMessages consuming ids', ids)
    this.signalrCore.invoke('MarkNotificationsAsRead', ids);
  }
}