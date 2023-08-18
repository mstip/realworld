import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';
import { SettingsComponent } from './pages/settings/settings.component';
import { EditorComponent } from './pages/editor/editor.component';
import { ArticleComponent } from './pages/article/article.component';
import { ProfileComponent } from './pages/profile/profile.component';

const routes: Routes = [
  {
    path: "",
    component: HomeComponent
  },
  {
    path: "login",
    component: LoginComponent
  },
  {
    path: "register",
    component: RegisterComponent
  },
  {
    path: "settings",
    component: SettingsComponent
  },
  {
    path: "editor",
    component: EditorComponent
  },
  {
    path: "editor/:slug",
    component: EditorComponent
  },
  {
    path: "article/:slug",
    component: ArticleComponent
  },
  {
    path: "profile/:username",
    component: ProfileComponent
  },
  {
    path: "profile/:username/favorites",
    component: ProfileComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
