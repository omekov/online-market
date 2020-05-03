import { NgModule } from '@angular/core';
import { Routes, RouterModule } from "@angular/router";
import { OrderComponent } from './order/order.component';
const routes: Routes = [
  {
    path: '',
    component: OrderComponent,
  }
]
@NgModule({
  declarations: [],
  imports: [
    RouterModule.forChild(routes),
  ],
  exports: [RouterModule]
})
export class OrderRoutingModule { }
