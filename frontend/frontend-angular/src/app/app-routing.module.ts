import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { NotLayoutComponent, LayoutComponent } from "./shared/layouts/layout.component";
import { NotfoundComponent } from './pages/notfound/notfound.component';

const routes: Routes = [
  {
    path: "",
    component: LayoutComponent,
    children: [
      {
        path: "order",
        loadChildren: () =>
          import("./pages/order/order.module").then((m) => m.OrderModule),
      },
      {
        path: "company",
        loadChildren: () =>
          import("./pages/company/company.module").then((m) => m.CompanyModule),
      },
      {
        path: "product",
        loadChildren: () =>
          import("./pages/product/product.module").then((m) => m.ProductModule),
      },
    ]
  },
  {
    path: "",
    component: NotLayoutComponent,
    children: [
      {
        path: "login",
        loadChildren: () =>
          import("./pages/login/login.module").then(m => m.LoginModule),
      },
      {
        path: "register",
        loadChildren: () =>
          import("./pages/register/register.module").then(
            m => m.RegisterModule
          ),
      },
    ],
  },
  {
    path: "**",
    redirectTo: "",
    component: NotfoundComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
