import { NgModule } from '@angular/core';
import { Routes, RouterModule } from "@angular/router";
import {CompanyComponent} from "./company/company.component";
const routes: Routes = [
  {
    path: '',
    component: CompanyComponent,
  }
]
@NgModule({
  declarations: [],
  imports: [
    RouterModule.forChild(routes),
  ],
  exports: [RouterModule]
})
export class CompanyRoutingModule { }
