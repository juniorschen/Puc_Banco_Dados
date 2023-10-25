import { Injectable } from '@angular/core'
import { HttpRequestOptions, Http } from '@nativescript/core';

import { Item } from './item'

@Injectable({
  providedIn: 'root',
})
export class ItemService {

  private getOptions(path: string, method: string, body = undefined): HttpRequestOptions {
    return {
      url: 'http://amos-junior.com:9999/' + path + "?Raw=true",
      method,
      headers: {
        'Content-Type': 'application/json'
      },
      content: body ? JSON.stringify(body) : undefined
    }
  }

  async getItems(): Promise<Array<Item>> {
    const options = this.getOptions("puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos", "GET");
    var result = await Http.request(options);
    if(result.statusCode != 200) {
      alert("Atenção um problema ocorreu durante a conexeção com serviço de dados, por favor contate o administrador do sistema repassando o seguinte erro: " + result.content.toString())
    }
    return result.content.toJSON();
  }

  async getItem(id: number): Promise<Item> {
    const options = this.getOptions(`puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos/${id}`, "GET");
    var result = await Http.request(options);
    if(result.statusCode != 200) {
      alert("Atenção um problema ocorreu durante a conexeção com serviço de dados, por favor contate o administrador do sistema repassando o seguinte erro: " + result.content.toString())
    }
    return result.content.toJSON();
  }

  async createItem(item) {
    const options = this.getOptions(`puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos`, "POST", item);
    var result = await Http.request(options);
    if(result.statusCode != 200) {
      alert("Atenção um problema ocorreu durante a conexeção com serviço de dados, por favor contate o administrador do sistema repassando o seguinte erro: " + result.content.toString())
    }
    return result.content.toJSON();
  }

  async updateItem(item) {
    const options = this.getOptions(`puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos/${item.Id}`, "PUT", item);
    var result = await Http.request(options);
    if(result.statusCode != 200) {
      alert("Atenção um problema ocorreu durante a conexeção com serviço de dados, por favor contate o administrador do sistema repassando o seguinte erro: " + result.content.toString())
    }
    return result.content.toJSON();
  }

  async deleteItem(id: number) {
    const options = this.getOptions(`puc/cursos/eng-software/disciplinas/bd/aulas/modelo-logico/entidades/robozinhos/${id}`, "DELETE");
    var result = await Http.request(options);
    if(result.statusCode != 200) {
      alert("Atenção um problema ocorreu durante a conexeção com serviço de dados, por favor contate o administrador do sistema repassando o seguinte erro: " + result.content.toString())
    }
    return result.content.toJSON();
  }
}
