import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
@Injectable({
  providedIn: 'root'
})
export class ApiService {
  baseUrl: string = 'http://localhost:5053/api/'
  constructor(private http: HttpClient) {
    this.baseUrl = environment.baseUrl
  }

  Create(url: string, body: any): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}${url}`, body)
  }
  GetAll(url: string): Observable<any[]> {
    return this.http.get<any[]>(`${this.baseUrl}${url}`)
  }
  GetById(url: string, id: number): Observable<any> {
    return this.http.get<any>(`${this.baseUrl}${url}/${id}`)
  }
  Update(url: string, id: number, body: any): Observable<any> {
    return this.http.put<any>(`${this.baseUrl}${url}/${id}`, body)
  }
  Delete(url: string, id: number): Observable<any> {
    return this.http.delete(`${this.baseUrl}${url}/${id}`)
  }
  Upload(url: string, formData: any): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}${url}`,formData, {
      reportProgress: true,
      observe: 'events'
    })
  }
}
