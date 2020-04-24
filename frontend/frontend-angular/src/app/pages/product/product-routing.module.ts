import { Routes, RouterModule } from "@angular/router";
import { ProductComponent } from './product/product.component';
import { NgModule } from '@angular/core';
const routes: Routes = [
  {
    path: '',
    component: ProductComponent,
  }
]
@NgModule({
  declarations: [],
  imports: [
    RouterModule.forChild(routes),
  ],
  exports: [RouterModule]
})
export class ProductRoutingModule { }
