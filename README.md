# IMAGE TAG HELPER BY GITHUB ACTION

## Example pipeline
```yaml
on:
   - push

jobs:
  deploy:
    runs-on: ["self-hosted"]
    steps:
      - uses: actions/checkout@master

      - name: 'Tag helper'
        uses: 'fphgov/actions-image-tag-helper@master'
        with:
          image_tag_prefix: 'prefix_'
          image_tag_suffix: '_suffix'
          is_add_short_commit_hash: true
    
      - name: 'Print short commit hash'
        run: echo ${{ env.SHORT_COMMIT_HASH }}
      
      - name: 'Print image tag'
        run: echo ${{ env.IMAGE_TAG }}
```

## Optional Arguments

| variable                           | description                         | required | default |
|------------------------------------|-------------------------------------|----------|---------|
| image_tag_prefix                   | Add prefix to image tag             | false    |         |
| image_tag_suffix                   | Add suffix to image tag             | false    |         |
| is_add_short_commit_hash           | Add short commit hash to image tag  | false    | false   |

## License

MIT License

Copyright (c) 2022 Municipality of Budapest

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.