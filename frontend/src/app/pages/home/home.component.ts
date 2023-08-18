import { Component, OnInit } from '@angular/core';
import { Articles } from 'src/app/models/articles.model';
import { Tags } from 'src/app/models/tags.model';
import { ArticleService } from 'src/app/services/article.service';
import { TagService } from 'src/app/services/tag.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  public tags: string[] = [];
  public articles: Articles = { articles: [], articlesCount: 0 };

  constructor(private tagService: TagService, private articleService: ArticleService) { }

  ngOnInit() {
    this.tagService.getTags().subscribe((data: Tags) => this.tags = data.tags, err => console.log(err));
    this.articleService.getArticles().subscribe((data: Articles) => this.articles = data, err => console.log(err));
  }
}
