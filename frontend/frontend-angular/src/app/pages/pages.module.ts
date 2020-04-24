import { NgModule, Component } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
@Component({
  selector: "pages",
  template: "<router-outlet></router-outlet>",
})
export class PagesComponent {}

const routes: Routes = [
  {
    path: "pages",
    component: PagesComponent,
    loadChildren: () => import('./product/product.module').then(m => m.ProductModule)
  },
];

@NgModule({
  declarations: [PagesComponent],
  imports: [RouterModule.forChild(routes)],
})
export class PagesModule {}
