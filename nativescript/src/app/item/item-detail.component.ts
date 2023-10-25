import { Component, NgZone, OnInit } from '@angular/core'
import { ActivatedRoute } from '@angular/router'
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RouterExtensions } from '@nativescript/angular';
import * as imagePickerPlugin from '@nativescript/imagepicker';

import { Item } from './item'
import { ItemService } from './item.service'
import { ImageAsset, ImageSource } from '@nativescript/core';

@Component({
  selector: 'ns-details',
  templateUrl: './item-detail.component.html',
})
export class ItemDetailComponent implements OnInit {

  private item: Item;

  public imageSrc: ImageSource;
  public thumbSize: number = 80;
  public previewSize: number = 300;
  public isSingleMode = true;
  public form: FormGroup;
  public isEditing = false;

  constructor(private itemService: ItemService, private route: ActivatedRoute, private fb: FormBuilder, private routerExtension: RouterExtensions,
    private _ngZone: NgZone) { }

  ngOnInit(): void {
    console.log(this.route.snapshot.params)
    const id = +this.route.snapshot.params.id;
    this.initializeForm();
    if (id) {
      this.isEditing = true;
      this.fillForm(id);
    }
  }


  private initializeForm() {
    this.form = this.fb.group({
      Id: [0],
      Nome: ["", [Validators.required]],
      Peso: [0, [Validators.required]],
    });
  }

  private async fillForm(id: number) {
    this.item = await this.itemService.getItem(id);
    this.form.patchValue(this.item);
    if (this.item.Foto) {
      this.imageSrc = await ImageSource.fromBase64(this.item.Foto);
    }
  }

  public async save() {
    if (this.form.invalid) {
      alert('Formulário inválido, preencha todos os campos!');
      return;
    }
    if (this.isEditing) {
      this.imageSrc.toBase64String
      await this.itemService.updateItem({
        Id: this.form.get('Id').value,
        Nome: this.form.get('Nome').value,
        Peso: Number(this.form.get('Peso').value),
        Foto: this.imageSrc ? this.imageSrc.toBase64String("png") : "",
      });
    } else {
      await this.itemService.createItem({
        Id: this.form.get('Id').value,
        Nome: this.form.get('Nome').value,
        Peso: Number(this.form.get('Peso').value),
        Foto: this.imageSrc ? this.imageSrc.toBase64String("png") : "",
      });
    }
    this.routerExtension.navigate(["/items"]);
  }

  public async delete() {
    await this.itemService.deleteItem(this.form.get('Id').value);
    this.routerExtension.navigate(["/items"]);
  }

  public onSelectSingleTap() {
    this.isSingleMode = true;

    let imagePicker = imagePickerPlugin.create({
      mode: 'single',
    });
    this.startSelection(imagePicker);
  }

  private startSelection(imagePicker) {
    imagePicker
      .authorize()
      .then(() => {
        this._ngZone.run(() => {
          this.imageSrc = null;
        });
        return imagePicker.present();
      })
      .then((selection) => {
        this._ngZone.run(() => {
          //console.log('Selection done: ' + JSON.stringify(selection));
          ImageSource.fromAsset(selection[0].asset).then(i => {
            this.imageSrc = i;
          });
        });
      })
      .catch(function (e) {
        console.log(e);
      });
  }
}