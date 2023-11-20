# RSS Parser
This is a package for RSS Parsing.

# Description
You can traverse the RSS feed array and access multiple elements.
You can extract thumbnails using HTML Meta Data.
It utilizes goroutines, which are management threads, and allows quick access to large amounts of data through parallel programming.
You can easily extract the elements below from that package:

- GUID
- Title
- Description
- Link
- Published

For additional information, please check <a href="https://github.com/mmcdole/gofeed">here</a>.

# Dependencies
- "github.com/mmcdole/gofeed"

# Get Started
1. Get the package.
```
go get github.com/dlworhd/rss-parser
```

2. Within an executable method such as the **main** function, write the code as shown below and run it.
```
func main() {

	wg := new(sync.WaitGroup)
	fp := gofeed.NewParser()

	rssUrls := []string{
		"https://netflixtechblog.com/feed",
	}

	wg.Add(len(rssUrls))

	var outerFeedArr [][]Feed

	for range rssUrls {
		outerFeedArr = append(outerFeedArr, nil)
	}

	for i, rss := range rssUrls {
		ProceedFeed(fp, rss, wg, outerFeedArr, i)
	}

	for _, innerFeedArr := range outerFeedArr {
		for _, feed := range innerFeedArr {
			fmt.Println("GUID: ", feed.guid)
			fmt.Println("Title: ", feed.title)
			fmt.Println("Thumbnail: ", feed.thumbnail)
			fmt.Println("Description : ", feed.description)
			fmt.Println("Link: ", feed.link)
			fmt.Println("Published: ", feed.published)
			fmt.Printf("\n")
		}
	}

	wg.Wait()

}
```
3. All feed information is stored in **outerFeedArr**, a two-dimensional array, and can be used as is.
![image](https://github.com/dlworhd/rss-parser/assets/102597172/c1652e0c-01ca-4286-ac1f-9bbd7b7a8002)


# Result
This is an example of execution results.

```
GUID:  https://medium.com/p/ac6483b52a51
Title:  Zero Configuration Service Mesh with On-Demand Cluster Discovery
Thumbnail:  https://miro.medium.com/v2/resize:fit:1000/0*xr1EBhghfcC8j4Uj
Description :  
Link:  https://netflixtechblog.com/zero-configuration-service-mesh-with-on-demand-cluster-discovery-ac6483b52a51?source=rss----2615bd06b42e---4
Published:  Tue, 29 Aug 2023 23:08:45 GMT

GUID:  https://medium.com/p/ed620b9c6225
Title:  Kubernetes And Kernel Panics
Thumbnail:  https://miro.medium.com/v2/resize:fit:1200/1*cjClRuyUQ67lu2shmjCObQ.png
Description :  
Link:  https://netflixtechblog.com/kubernetes-and-kernel-panics-ed620b9c6225?source=rss----2615bd06b42e---4
Published:  Fri, 27 Oct 2023 16:05:58 GMT

GUID:  https://medium.com/p/0d83f5a00d08
Title:  Streaming SQL in Data Mesh
Thumbnail:  https://miro.medium.com/v2/resize:fit:1200/1*VxZlXPDb8-d7Xf4kfSulnw.png
Description :  
Link:  https://netflixtechblog.com/streaming-sql-in-data-mesh-0d83f5a00d08?source=rss----2615bd06b42e---4
Published:  Fri, 03 Nov 2023 21:48:50 GMT

GUID:  https://medium.com/p/f68830617dd1
Title:  1. Streamlining Membership Data Engineering at Netflix with Psyberg
Thumbnail:  https://miro.medium.com/v2/resize:fit:1200/0*i3Q9OtyFGyxh0Zon
Description :  
Link:  https://netflixtechblog.com/1-streamlining-membership-data-engineering-at-netflix-with-psyberg-f68830617dd1?source=rss----2615bd06b42e---4
Published:  Wed, 15 Nov 2023 03:24:49 GMT

GUID:  https://medium.com/p/afd64e6a5bf8
Title:  Detecting Speech and Music in Audio Content
Thumbnail:  https://miro.medium.com/v2/da:true/resize:fit:720/1*W-4QTtWN_NDQt4HWr3EBBQ.gif
Description :  
Link:  https://netflixtechblog.com/detecting-speech-and-music-in-audio-content-afd64e6a5bf8?source=rss----2615bd06b42e---4
Published:  Tue, 14 Nov 2023 01:55:45 GMT

GUID:  https://medium.com/p/936766f0017c
Title:  Building In-Video Search
Thumbnail:  https://miro.medium.com/v2/resize:fit:1000/0*32RfnKGMENXaqEX8
Description :  
Link:  https://netflixtechblog.com/building-in-video-search-936766f0017c?source=rss----2615bd06b42e---4
Published:  Mon, 06 Nov 2023 17:35:19 GMT

GUID:  https://medium.com/p/1d273b3aaefb
Title:  Diving Deeper into Psyberg: Stateless vs Stateful Data Processing
Thumbnail:  https://miro.medium.com/v2/resize:fit:389/1*RnFDv0pCKpSxEiWBF_e2kw.png
Description :  
Link:  https://netflixtechblog.com/2-diving-deeper-into-psyberg-stateless-vs-stateful-data-processing-1d273b3aaefb?source=rss----2615bd06b42e---4
Published:  Wed, 15 Nov 2023 03:25:13 GMT

GUID:  https://medium.com/p/b8cd145491cc
Title:  AVA Discovery View: Surfacing Authentic Moments
Thumbnail:  https://miro.medium.com/v2/resize:fit:890/0*o3U-Eg_XXmXzmMMO
Description :  
Link:  https://netflixtechblog.com/ava-discovery-view-surfacing-authentic-moments-b8cd145491cc?source=rss----2615bd06b42e---4
Published:  Thu, 17 Aug 2023 22:07:14 GMT

GUID:  https://medium.com/p/260fbe366fe2
Title:  Psyberg: Automated end to end catch up
Thumbnail:  https://miro.medium.com/v2/resize:fit:1200/1*NhFne1sCHTTW8ZzahSZqAQ.png
Description :  
Link:  https://netflixtechblog.com/3-psyberg-automated-end-to-end-catch-up-260fbe366fe2?source=rss----2615bd06b42e---4
Published:  Wed, 15 Nov 2023 03:25:23 GMT

GUID:  https://medium.com/p/4dc4ce2011ef
Title:  The Next Step in Personalization: Dynamic Sizzles
Thumbnail:  https://miro.medium.com/v2/resize:fit:948/0*OxP3aIMkQW6d-BjK
Description :  
Link:  https://netflixtechblog.com/the-next-step-in-personalization-dynamic-sizzles-4dc4ce2011ef?source=rss----2615bd06b42e---4
Published:  Wed, 08 Nov 2023 20:56:53 GMT
```

