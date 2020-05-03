import { Component, OnInit, ViewChild, ElementRef } from "@angular/core";
import { ApiService } from "../../../services/api.service";
import { IProduct, ICategory } from "src/app/interfaces/product";
import { HttpEventType, HttpErrorResponse } from "@angular/common/http";
import { of } from "rxjs";
import { catchError, map } from "rxjs/operators";
@Component({
  selector: "app-order",
  templateUrl: "./order.component.html",
  styleUrls: ["./order.component.scss"],
})
export class OrderComponent implements OnInit {
  @ViewChild("fileUpload", { static: false }) fileUpload: ElementRef;
  files = [];
  constructor(private apiService: ApiService) {}

  products: IProduct[] = [];
  categories: ICategory[] = [];
  ngOnInit(): void {
    this.apiService
      .GetAll("products")
      .subscribe((data) => (this.products = data));
    this.apiService
      .GetAll("categories")
      .subscribe((data) => (this.categories = data));
  }
  filterOrder(id: number): void {
    return;
  }
  addCart() {}
  uploadPhoto(file) {
    const formData = new FormData();
    formData.append("file", file.data);
    file.inProgress = true;
    this.apiService
      .Upload("uploadproductphoto", formData)
      .pipe(
        map((event) => {
          switch (event.type) {
            case HttpEventType.UploadProgress:
              file.progress = Math.round((event.loaded * 100) / event.total);
              break;
            case HttpEventType.Response:
              return event;
          }
        }),
        catchError((error: HttpErrorResponse) => {
          file.inProgress = false;
          return of(`${file.data.name} upload failed.`);
        })
      )
      .subscribe((event: any) => {
        if (typeof event === "object") {
          console.log(event.body);
        }
      });
  }
  private uploadFiles() {
    this.fileUpload.nativeElement.value = "";
    this.files.forEach((file) => {
      this.uploadPhoto(file);
    });
  }
  onClick() {
    const fileUpload = this.fileUpload.nativeElement;
    fileUpload.onchange = () => {
      for (let index = 0; index < fileUpload.files.length; index++) {
        const file = fileUpload.files[index];
        this.files.push({ data: file, inProgress: false, progress: 0 });
      }
      this.uploadFiles();
    };
    fileUpload.click();
  }
}
