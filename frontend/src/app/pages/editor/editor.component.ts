import { Component } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.css']
})
export class EditorComponent {
  public errors : string[] = [];


  editorForm = this.fb.group({
    title: ['', Validators.compose([Validators.required, Validators.minLength(5)])],
    description: ['', Validators.compose([Validators.required, Validators.minLength(5)])],
    body: ['', Validators.compose([Validators.required, Validators.minLength(5)])],
    tags: ['']
  })

  constructor(private fb: FormBuilder) {

  }

  onPublish() {
    this.errors = [];
    // console.log(this.editorForm.value)
    // console.log(this.editorForm.status)
    // console.log(this.editorForm.valid)
    // console.log(this.editorForm.get('title')?.errors)
    console.log(this.editorForm.get('title')?.errors);
    if(!this.editorForm.valid) {
      if(this.editorForm.get('title')?.errors !== null) {
        this.errors.push("title must be atleast 5 characters long")
      }
      if(this.editorForm.get('description')?.errors !== null) {
        this.errors.push("description must be atleast 5 characters long")
      }
      if(this.editorForm.get('body')?.errors !== null) {
        this.errors.push("body must be atleast 5 characters long")
      }
    }
  }

}
