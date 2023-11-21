# RSS Parser
This is a package for RSS Parsing.

# Description
Traverse the RSS feed array and access multiple elements.<br>
Extract thumbnails using HTML Meta Data.<br>
It utilizes goroutines, which are management threads, and allows quick access to large amounts of data through parallel programming.<br>

Extract easily the elements below from that package:

- GUID
- Title
- Description
- Link
- Published

# Applied
- GoRoutine
- Feed Struct 
- Return 2D Array

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
		rss_parser.ProceedFeed(fp, rss, wg, outerFeedArr, i)
	}

	for _, innerFeedArr := range outerFeedArr {
		for _, feed := range innerFeedArr {
			fmt.Println("GUID: ", feed.guid)
			fmt.Println("Title: ", feed.title)
			fmt.Println("Thumbnail: ", feed.thumbnail)
			fmt.Println("Description : ", feed.description)
			fmt.Println("Link: ", feed.link)
			fmt.Println("Published: ", feed.published)
		}
	}

	wg.Wait()

}
```
3. All feed information is stored in **outerFeedArr**, a two-dimensional array, and can be used as is.
<img width=100% src="https://github.com/dlworhd/rss-parser/assets/102597172/b0ee14ce-2a61-4c0a-87bf-0a25317f67e1">


# Result
This is an example of execution results.

```
GUID:  https://toss.tech/article/how-to-learn-new-domain
Title:  디자이너가 새로운 도메인을 빠르게 학습하는 법
Thumbnail:  https://toss.tech/wp-content/uploads/2023/03/og-image-1536x768.png
Description :  토스는 일이 굉장히 빠르게 돌아가는 조직이에요. 저는 1년 반 동안 무려 4개의 새로운 분야를 학습해야 했어요. 이때 제가 쓴 방법들을 알려드릴게요.
Link:  https://toss.tech/article/how-to-learn-new-domain
Published:  Tue, 07 Mar 2023 05:56:37 GMT

GUID:  https://toss.tech/article/1st_ux_research
Title:  리서치를 하고 싶어하는 사람을 리서치하세요
Thumbnail:  https://static.toss.im/assets/toss-tech/research_og.png
Description :  입사하자마자 사용자가 아닌 동료들부터 인터뷰했던 이유. 토스의 첫 UX 리서처로 자리잡아 갔던 과정을 소개할게요.
Link:  https://toss.tech/article/1st_ux_research
Published:  Thu, 27 Apr 2023 11:25:00 GMT

GUID:  https://toss.tech/article/recommend-just-one
Title:  추천할 때는 제일 좋은 것 하나면 된다
Thumbnail:  https://static.toss.im/assets/tech-blog/thumbnail-02.png?v=2
Description :  제품을 만들 때, 사용자의 고민을 덜어주는 것이 중요해요. 사용자에게는 어떤 선택지를 주는 것이 좋을까요?
Link:  https://toss.tech/article/recommend-just-one
Published:  Thu, 23 Mar 2023 07:43:33 GMT

GUID:  https://toss.tech/article/improving-code-quality-via-eslint-and-ast
Title:  ESLint와 AST로 코드 퀄리티 높이기
Thumbnail:  https://og.toss.tech?title=ESLint%EC%99%80%20AST%EB%A1%9C%0D%0A%EC%BD%94%EB%93%9C%20%ED%80%84%EB%A6%AC%ED%8B%B0%20%EB%86%92%EC%9D%B4%EA%B8%B0&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Ftoss-tech%2Fright.png&v=2
Description :  ESLint와 AST로 토스에서 코드 퀄리티를 높인 방법에 대해 소개드려요.
Link:  https://toss.tech/article/improving-code-quality-via-eslint-and-ast
Published:  Fri, 31 Mar 2023 06:47:53 GMT

GUID:  https://toss.tech/article/engineering-note-1
Title:  웹에서 복잡한 퍼널 쉽게 관리하기
Thumbnail:  https://static.toss.im/assets/payments/contents/engineering_thumb.jpg
Description :  토스페이먼츠 프론트엔드 챕터에서 웹에서 퍼널을 손쉽게 관리하기 위해 했던 고민과 해결 방법을 공유해요.
Link:  https://toss.tech/article/engineering-note-1
Published:  Wed, 18 Oct 2023 00:00:00 GMT

GUID:  https://toss.tech/article/slash23-data
Title:  대규모 로그 처리도 OK! Elasticsearch 클러스터 개선기
Thumbnail:  https://static.toss.im/career-resource/techblog_slash23_og_04_이준환.png
Description :  일평균 56억 건 이상 수집되는 토스증권의 로그, 어떻게 효율적으로 처리할까요?
큰 폭으로 늘어나는 대규모 로그 처리를 위한 ‘Elasticsearch 클러스터 개선 경험’을 소개합니다.
Link:  https://toss.tech/article/slash23-data
Published:  Thu, 12 Oct 2023 07:03:00 GMT

GUID:  https://toss.tech/article/only-3-users-product
Title:  사용자 3명인 제품, 만들어야 할까?
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-12.png?v=2
Description :  사용자가 3명 뿐인 제품, 쓰는 사람이 적으면 개선을 하지 말아야 할까요? 토스 콘텐츠 시스템(TCS)을 만들면서, 최소한의 비용으로 최대한의 효율을 내기 위해 고민했던 지점들을 소개해드려요.
Link:  https://toss.tech/article/only-3-users-product
Published:  Thu, 22 Dec 2022 07:20:35 GMT

GUID:  https://toss.tech/article/nodejs-security-contribution
Title:  Node.js url.parse() 취약점 컨트리뷰션
Thumbnail:  https://static.toss.im/assets/toss-tech/node%20js-security.png
Description :  토스 보안기술팀은 안전한 금융 서비스를 제공하기 위한 연구를 수행하고 있어요.
많은 서비스에서 사용되고 있는 Node.js의 취약점을 분석하고 안전하게 패치될 수 있도록 기여했던 과정을 소개드려요.
Link:  https://toss.tech/article/nodejs-security-contribution
Published:  Fri, 12 May 2023 08:00:00 GMT

GUID:  https://toss.tech/article/isomorphic-javascript
Title:  환경 고민없이 개발하기
Thumbnail:  https://static.toss.im/illusts-common/isomorphic-tech-blog-thumb-nail.png
Description :  서버 사이드 렌더링 작동 방식과 Isomorphic에 대해 소개드려요.
Link:  https://toss.tech/article/isomorphic-javascript
Published:  Fri, 01 Sep 2023 03:10:00 GMT

GUID:  https://toss.tech/article/smart-polyfills
Title:  똑똑하게 브라우저 Polyfill 관리하기
Thumbnail:  https://og.toss.tech?title=%EB%98%91%EB%98%91%ED%95%98%EA%B2%8C%20%EB%B8%8C%EB%9D%BC%EC%9A%B0%EC%A0%80%0D%0APolyfill%20%EA%B4%80%EB%A6%AC%ED%95%98%EA%B8%B0&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Ftoss-tech%2Fjs_right.png&v=2
Description :  현대적인 JavaScript를 쓰면서도 넓은 범위의 기기를 지원하기 위한 Polyfill을 어떻게 똑똑하게 설정할 수 있는지 소개합니다.
Link:  https://toss.tech/article/smart-polyfills
Published:  Sat, 21 Jan 2023 03:52:16 GMT

GUID:  https://toss.tech/article/smart-web-service-cache
Title:  웹 서비스 캐시 똑똑하게 다루기
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-02-webcash-og.png?v=2
Description :  웹 성능을 위해 꼭 필요한 캐시, 제대로 설정하기 쉽지 않습니다. 토스 프론트엔드 챕터에서 올바르게 캐시를 설정하기 위한 노하우를 공유합니다.
Link:  https://toss.tech/article/smart-web-service-cache
Published:  Thu, 29 Apr 2021 08:08:00 GMT

GUID:  https://toss.tech/article/how-to-manage-test-dependency-in-gradle
Title:  테스트 의존성 관리로 높은 품질의 테스트 코드 유지하기
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-10-test-code-og.png?v=2
Description :  혹시 테스트 코드에서도 의존성을 관리해본 적이 있으실까요? 해당 포스트에서는 Gradle의 java-test-fixtures 플러그인을 사용하여 테스트 의존성 관리를 통해 높은 품질의 테스트 코드를 유지하는 방법을 알아봅니다.
Link:  https://toss.tech/article/how-to-manage-test-dependency-in-gradle
Published:  Wed, 08 Jun 2022 15:17:10 GMT

GUID:  https://toss.tech/article/dont-trust-your-gut
Title:  내 아이디어를 너무 믿지 마세요
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-14-og.png?v=2
Description :  너무 좋은 아이디어라고 생각해서 만들었는데 실제 유저들의 반응은 정반대였어요.
Link:  https://toss.tech/article/dont-trust-your-gut
Published:  Tue, 28 Mar 2023 08:31:40 GMT

GUID:  https://toss.tech/article/value-first-cost-later
Title:  토스 디자인 원칙 Value first, Cost later
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/%E1%84%87%E1%85%A9%E1%86%AB%E1%84%86%E1%85%AE%E1%86%AB%20%E1%84%8A%E1%85%A5%E1%86%B7%E1%84%82%E1%85%A6%E1%84%8B%E1%85%B5%E1%86%AF_%E1%84%87%E1%85%A2%E1%86%AF%E1%84%85%E1%85%B2%E1%84%91%E1%85%A5%E1%84%89%E1%85%B3%E1%84%90%E1%85%B3.png?v=2
Description :  토스의 제품 디자인 원칙 중에는 Value first, cost later라는 항목이 있어요. 비용을 말하기 전에 가치를 명확하게 전달해야한다는 내용이에요. 가치를 먼저 보여주는 것 만으로도 지표를 눈에 띄게 개선한 사례를 소개해드릴게요.
Link:  https://toss.tech/article/value-first-cost-later
Published:  Thu, 16 Mar 2023 09:49:46 GMT

GUID:  https://toss.tech/article/2022-product-designer-challenge
Title:  디자이너의, 디자이너에 의한, 디자이너를 위한 채용 설계하기
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-3-og.png?v=2
Description :  디자이너가 디자이너 채용을 기획한다면 어떻게 될까요? 프로덕트 디자이너 챌린지는 지원자인 디자이너의 입장에서 기획한 채용이에요. 회사가 어떻게 지원자를 검증할지만 생각한 게 아니라, 어떻게 하면 지원자의 역량을 잘 끌어낼 수 있을지 생각하며 3가지 방법을 선택했어요.
Link:  https://toss.tech/article/2022-product-designer-challenge
Published:  Thu, 22 Sep 2022 06:55:39 GMT

GUID:  https://toss.tech/article/how-platform-designer-make-effectiveness
Title:  플랫폼 디자이너가 효율을 만들어내는 법
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-9-og.png?v=2
Description :  수많은 가짓수들을 단순히 더하기로 추가하는 것이 아니라 곱하기로 확장하면서 극단적인 효율과 심미성까지 챙겼던 과정들을 플로우차트 제작 프로젝트를 예시로 소개해드릴게요.
Link:  https://toss.tech/article/how-platform-designer-make-effectiveness
Published:  Thu, 08 Dec 2022 09:40:30 GMT

GUID:  https://toss.tech/article/faster-initial-rendering
Title:  조금만 신경써서 초기 렌더링 빠르게 하기 (feat. JAM Stack)
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-07-jam-stack-og.png?v=2
Description :  SPA(Single Page Application) 구조로 웹 프론트엔드 애플리케이션이 개발되면서 초기 렌더링 속도는 프런트엔드 개발자에게 중요한 과제 중 하나가 되었습니다. 사용자 경험에 영향을 줄 수 있는 가장 큰 요소 중 하나가 바로 속도이기 때문입니다.
Link:  https://toss.tech/article/faster-initial-rendering
Published:  Wed, 09 Feb 2022 05:25:46 GMT

GUID:  https://toss.tech/article/8-writing-principles-of-toss
Title:  토스의 8가지 라이팅 원칙들
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-8-og.png?v=2
Description :  토스의 문구는 8가지 라이팅 원칙을 고려하면서 쓰고 있어요. 사람이 말하는 것 같은 문장을 지향하면서요.
Link:  https://toss.tech/article/8-writing-principles-of-toss
Published:  Tue, 15 Nov 2022 08:53:48 GMT

GUID:  https://toss.tech/article/product-branding-team
Title:  프로덕트 브랜딩, 어떻게 시작해야할까?
Thumbnail:  https://static.toss.im/assets/toss-tech/og-1st-product-brand-designer.png
Description :  처음 프로덕트 브랜딩 팀이 만들어졌을 때, 어떻게 팀의 정체성을 찾아나갔는지 공유할게요.
Link:  https://toss.tech/article/product-branding-team
Published:  Thu, 10 Aug 2023 09:15:00 GMT

GUID:  https://toss.tech/article/1st_uxwriter
Title:  첫 UX writer는 무슨 일을 해야 할까
Thumbnail:  https://static.toss.im/assets/toss-tech/og_1st_uxwriter.png
Description :  이게 정말 UX writer만 할 수 있는 일일까? 내가 해야 하는 일이 이것일까? 생각했을 때 확신이 없었어요. 일단은 모든 업무를 가리지 않고 다 받아서 무작정 열심히 처리했었죠.
Link:  https://toss.tech/article/1st_uxwriter
Published:  Fri, 14 Apr 2023 03:23:00 GMT

GUID:  https://toss.tech/article/touch-and-turn-tossbankcard
Title:  직접 만지고, 돌리는 토스뱅크카드 인터랙션
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/DesignTech-og.png?v=2
Description :  토스뱅크카드의 중요한 디자인 컨셉은 앞면과 뒷면의 색상이 다르다는 것인데요. 지금까지는 이 정보를 표현하기 위해서 이미지를 두 장 쓰거나, 영상을 만들었죠. 하지만 그걸로는 부족했어요. 저는 모바일 화면에서도 실물 카드를 보는 것과 똑같은 경험을 만들고 싶었어요.
Link:  https://toss.tech/article/touch-and-turn-tossbankcard
Published:  Thu, 24 Nov 2022 10:16:25 GMT

GUID:  https://toss.tech/article/how-to-write-error-message
Title:  좋은 에러 메시지를 만드는 6가지 원칙
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-2-og.png?v=2
Description :  좋은 에러 메시지란, 사용자가 다음 단계로 갈 수 있게 돕는 메시지예요. 사용자가 다음 단계로 가기 위해 필요한 내용은 어떤 것들이 있을까요? 토스는 6가지 에러 메시지 원칙을 생각하며 문구를 쓰고 있어요.
Link:  https://toss.tech/article/how-to-write-error-message
Published:  Wed, 21 Sep 2022 14:24:51 GMT

GUID:  https://toss.tech/article/slash23-iOS
Title:  레고처럼 조립하는 토스 앱
Thumbnail:  https://static.toss.im/assets/toss-tech/slash_juneseokbeomgun_og.png
Description :  수많은 서비스를 담고 있는 대규모 iOS 앱에 어울리는 아키텍처는 무엇일까요?
프로젝트 간의 의존성과 모듈 간의 결합도를 낮춰, 더 효율적인 서비스 개발, 관리를 이뤄낸 과정을 소개합니다. 
Link:  https://toss.tech/article/slash23-iOS
Published:  Tue, 22 Aug 2023 08:46:00 GMT

GUID:  https://toss.tech/article/product-designer-challenge-review
Title:  경험이 부족한 나도 토스에 지원해도 될까?
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-4-og.png?v=2
Description :  이력서, 포트폴리오가 아닌 과제 중심으로 역량을 평가한 프로덕트 디자이너 챌린지. 신입이라 선뜻 지원을 못했었는데, 챌린지를 통해 입사할 수 있었어요. 챌린지에 참여하게 된 계기와 문제를 해결하면서 고민했던 부분들, 장표를 어떻게 구성했는지까지 자세한 후기를 공유해드릴게요.
Link:  https://toss.tech/article/product-designer-challenge-review
Published:  Tue, 04 Oct 2022 12:13:05 GMT

GUID:  https://toss.tech/article/tosspayments-restdocs
Title:  tosspayments-restdocs: 선언형 문서 작성 라이브러리
Thumbnail:  https://toss.tech/wp-content/uploads/2023/03/00017-3291509353.png?v=2
Description :  REST Docs 를 최소한의 코드로 작성하면서 변화에도 더 유연하게 대처할 수 있는 tosspayments-restdocs 라이브러리와, 라이브러리에 녹인 기술들을 소개합니다.
Link:  https://toss.tech/article/tosspayments-restdocs
Published:  Wed, 22 Mar 2023 05:29:51 GMT

GUID:  https://toss.tech/article/kotlin-dsl-restdocs
Title:  Kotlin으로 DSL 만들기: 반복적이고 지루한 REST Docs 벗어나기
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-08-rest-docs-og.png?v=2
Description :  토스페이먼츠에서는 API docs를 REST Docs를 사용해서 작성할 수 있도록 권장하고 있습니다. 이 글에서는 DSL을 통해서 반복적인 REST Docs 테스트 코드 작성을 줄일 수 있는 방법을 알아봅니다.
Link:  https://toss.tech/article/kotlin-dsl-restdocs
Published:  Sun, 10 Apr 2022 20:24:00 GMT

GUID:  https://toss.tech/article/how-to-work-health-check-in-spring-boot-actuator
Title:  Spring Boot Actuator의 헬스체크 살펴보기
Thumbnail:  https://static.toss.im/illusts-content/img-tech-cover.png
Description :  서버의 상태를 알려주는 헬스 체크에 대해 알고 계시나요? 단순히 200 OK만 내려주겠거니 하고 별로 신경을 안 쓰고 계셨나요? 해당 포스트에서는 Spring Boot Actuaor가 제공해주는 헬스 체크는 어떤 식으로 서버의 상태를 점검하는지, 어떤 부분을 주의하며 사용해야하는지 알아봅니다.
Link:  https://toss.tech/article/how-to-work-health-check-in-spring-boot-actuator
Published:  Sat, 01 Apr 2023 06:38:57 GMT

GUID:  https://toss.tech/article/slash23-corebanking
Title:  은행 최초 코어뱅킹 MSA 전환기 (feat. 지금 이자 받기)
Thumbnail:  https://static.toss.im/career-resource/techblog_slash23_og_02.png
Description :  수십 년간 정체되어 있던 전통적인 은행 시스템의 모놀리식 소프트웨어 아키텍처를 MSA로 전환할 수 있을까요? 
토스뱅크의 ‘코어뱅킹 MSA 전환’ 사례를 통해 향후 은행 시스템이 나아가야 할 방향을 소개합니다.
Link:  https://toss.tech/article/slash23-corebanking
Published:  Thu, 31 Aug 2023 14:37:00 GMT

GUID:  https://toss.tech/article/template-literal-types
Title:  Template Literal Types로 타입 안전하게 코딩하기
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-05-template-literal-og.png?v=2
Description :  TypeScript 코드베이스의 타입 안전성을 한 단계 올려줄 수 있는 Template Literal Type의 뜻과 응용에 대해 알아봅니다.
Link:  https://toss.tech/article/template-literal-types
Published:  Fri, 14 May 2021 06:26:37 GMT

GUID:  https://toss.tech/article/nestjs-custom-decorator
Title:  NestJS 환경에 맞는 Custom Decorator 만들기
Thumbnail:  https://static.toss.im/assets/finance-tips/img-financetip-og2.png
Description :  NestJS에서 데코레이터를 만들기 위해서는 NestJS의 DI와 메타 프로그래밍 환경 등을 고려해야 합니다. 어떻게 하면 이러한 NestJS 환경에 맞는 데코레이터를 만들 수 있을지 고민해보았습니다.
Link:  https://toss.tech/article/nestjs-custom-decorator
Published:  Mon, 21 Nov 2022 15:03:49 GMT

GUID:  https://toss.tech/article/tossinvest-qa-integration-test
Title:  토스증권 QA 문화 ‘통합테스트’를 아시나요? (feat. 해외주식)
Thumbnail:  https://og.toss.tech?title=%ED%86%A0%EC%8A%A4%EC%A6%9D%EA%B6%8C%EC%9D%98%20QA%20%EB%AC%B8%ED%99%94%3A%0D%0A%E2%80%98%ED%86%B5%ED%95%A9%20%ED%85%8C%EC%8A%A4%ED%8A%B8%E2%80%99&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Ftech-blog%2Fog%20image%2Fthumnail%2F01-og.png&v=2
Description :  토스증권 해외주식 출시 전에 사내 임직원 대상으로 진행한 ‘통합테스트’에 대해 소개합니다. 통합테스트 진행 방식을 참고하여 간단한 규칙과 사용자 시나리오를 활용해 사용자의 반응을 미리 확인해 보세요.
Link:  https://toss.tech/article/tossinvest-qa-integration-test
Published:  Mon, 12 Dec 2022 12:01:48 GMT

GUID:  https://toss.tech/article/interaction
Title:  인터랙션, 꼭 넣어야 해요?
Thumbnail:  https://static.toss.im/simplicity23/graphics/sim23-interaction-og.png
Description :  빠른 속도를 중요시하는 토스에서 어떻게 팀원들을 인터랙션에 공감하게 하고 슬릭한 제품을 만들어나갈 수 있었는지 소개할게요.
Link:  https://toss.tech/article/interaction
Published:  Thu, 07 Sep 2023 05:27:00 GMT

GUID:  https://toss.tech/article/simplicity23
Title:  Simplicity23, 오늘도 문제를 해결하고 있을 모든 디자이너에게
Thumbnail:  https://static.toss.im/career-resource/article-cover.png
Description :  토스 디자인 컨퍼런스 사전 신청 오픈 (5.15~5.21)
Link:  https://toss.tech/article/simplicity23
Published:  Mon, 15 May 2023 03:00:00 GMT

GUID:  https://toss.tech/article/frontend-diving-club-agora
Title:  프론트엔드 다이빙클럽에서 만나는 아고라: 다른 회사에선 테스트 코드 어떻게 짜요?
Thumbnail:  https://static.toss.im/ml-illust/img-people-check-outfut.jpg
Description :  ‘프론트엔드 테스팅’을 주제로 진행된 프론트엔드 다이빙 클럽의 네번째 모임을 공유합니다.
Link:  https://toss.tech/article/frontend-diving-club-agora
Published:  Wed, 11 Oct 2023 02:39:00 GMT

GUID:  https://toss.tech/article/engineering-note-2
Title:  null 리턴은 왜 나쁠까?
Thumbnail:  https://static.toss.im/ml-illust/_20231108_104549157.png
Description :  코드 복잡성 관리 측면에서 의미를 축약한 표현의 문제와 해결 방법을 예제로 알아봐요.
Link:  https://toss.tech/article/engineering-note-2
Published:  Wed, 08 Nov 2023 01:00:00 GMT

GUID:  https://toss.tech/article/slash23-server
Title:  토스는 Gateway 이렇게 씁니다
Thumbnail:  https://static.toss.im/career-resource/techblog_slash23_og_05_최준우.png
Description :  더 안전하고 안정적인 서비스 운영을 위해서 ‘gateway’를 어떻게 사용해야 할까요? 
토스의 수많은 마이크로서비스 로직을 공통화하기 위한 gateway 운영 방법에 대해 소개합니다.
Link:  https://toss.tech/article/slash23-server
Published:  Thu, 12 Oct 2023 10:52:00 GMT

GUID:  https://toss.tech/article/frontend-declarative-code
Title:  선언적인 코드 작성하기
Thumbnail:  https://og.toss.tech?title=%EC%84%A0%EC%96%B8%EC%A0%81%EC%9D%B8%20%EC%BD%94%EB%93%9C%0D%0A%EC%9E%91%EC%84%B1%ED%95%98%EA%B8%B0&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Ftoss-tech%2Fdeclarative_right.png&v=2
Description :  선언적인 코드, 토스 프론트엔드 챕터는 어떻게 생각을 하고 있을까요?
Link:  https://toss.tech/article/frontend-declarative-code
Published:  Thu, 16 Mar 2023 12:41:13 GMT

GUID:  https://toss.tech/article/kotlin-result
Title:  에러 핸들링을 다른 클래스에게 위임하기 (Kotlin 100% 활용)
Thumbnail:  https://static.toss.im/assets/tech-blog/og%20image/thumnail/error.png?v=2
Description :  Kotlin의 Result로 MSA에서 에러가 전파되지 않는 안전한 환경을 만드는 방법을 알아봅니다.
Link:  https://toss.tech/article/kotlin-result
Published:  Sat, 14 May 2022 14:38:16 GMT

GUID:  https://toss.tech/article/designer-onboarding
Title:  토스에 처음 입사한 디자이너를 위한 온보딩 프로그램
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-1201-og.png?v=2
Description :  토스 디자인 챕터는 온보딩을 위한 다양한 프로그램이 있어요. 디자이너가 직접 설계한 디자이너 온보딩, 어떤 것들을 신경 썼는지 하나씩 소개해드릴게요.
Link:  https://toss.tech/article/designer-onboarding
Published:  Thu, 01 Dec 2022 06:31:41 GMT

GUID:  https://toss.tech/article/1st_interaction_designer
Title:  첫 인터랙션 디자이너가 문제를 해결하는 법
Thumbnail:  https://static.toss.im/assets/toss-tech/thumnail-1st-interactiondesigner.png
Description :  처음 입사했을 때 인터랙션 디자이너의 역할에 대해 막연하게 느꼈을 때가 있었어요. 하지만 제품의 문제를 하나씩 해결하면서 저의 역할을 정의해나갈 수 있었어요. 어떤 시행착오가 있었는지 공유해 드릴게요.
Link:  https://toss.tech/article/1st_interaction_designer
Published:  Thu, 04 May 2023 10:48:00 GMT

GUID:  https://toss.tech/article/uss-card
Title:  30대 디자이너가 10대 전용 카드를 만든다면?
Thumbnail:  https://static.toss.im/assets/toss-tech/sim23-uss-og.png
Description :  '요즘 10대는 어떤 카드를 좋아할까?', '30대인 내가 10대 친구들이 좋아하는 디자인을 할 수 있을까?' 청소년 전용 카드를 만들며 맞닥뜨렸던 혹독한 진실과, 그 과정에서 얻은 인사이트를 공유하고 싶어요.
Link:  https://toss.tech/article/uss-card
Published:  Wed, 23 Aug 2023 06:25:00 GMT

GUID:  https://toss.tech/article/thinking-user-perspective
Title:  우리에게 당연하지만 사용자는 아닌 것들
Thumbnail:  https://static.tosspayments.com/public/permanent/service/contents/tech-og.png?v=2
Description :  사용자 관점에서 본다는 게 엄청난 게 아닌데도, 왠지 엄청난 것으로 개선해야 할 것 같은 느낌이 들 때도 있죠. 그래서 오늘은 사용자 관점을 활용해 개선한 굉장히 가벼운 사례들을 가져와봤어요.
Link:  https://toss.tech/article/thinking-user-perspective
Published:  Thu, 19 Jan 2023 09:15:32 GMT

GUID:  https://toss.tech/article/22205
Title:  토스의 이모지 폰트, 토스페이스 제작기
Thumbnail:  https://static.toss.im/3d/tossface-part-thumbnail.png
Description :  세계적인 IT 기업에서나 만드는 이모지 폰트를 어떻게 한국 금융 플랫폼 토스에서 만들게 됐을까? 토스의 이모지 폰트, 토스페이스 제작의 모든 것을 공개합니다.
Link:  https://toss.tech/article/22205
Published:  Mon, 17 Jul 2023 03:00:00 GMT

GUID:  https://toss.tech/article/4-ways-for-minimum-input
Title:  2초만에 불필요한 클릭 없애는 4가지 방법
Thumbnail:  https://og.toss.tech?title=2%EC%B4%88%EB%A7%8C%EC%97%90%20%EB%B6%88%ED%95%84%EC%9A%94%ED%95%9C%20%ED%81%B4%EB%A6%AD%20%EC%97%86%EC%95%A0%EB%8A%94%204%EA%B0%80%EC%A7%80%20%EB%B0%A9%EB%B2%95&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Fhomepage%2Ftosstech%2Fog%2Fdesign-article-5-2-og.png&v=2
Description :  토스 앱에는 little big detail 이 많은데요, 그중 하나를 소개하고 싶어요.
바로 불필요한 클릭 없애는 방법이에요.
Link:  https://toss.tech/article/4-ways-for-minimum-input
Published:  Thu, 20 Oct 2022 02:20:00 GMT

GUID:  https://toss.tech/article/toss-signup-process
Title:  거꾸로 입력하는 가입 화면, 처음에 어떻게 떠올렸을까?
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-1-og.png?v=2
Description :  토스의 회원 가입 화면에선 스크롤을 내릴 필요가 없어요. 필요한 정보들을 거꾸로 입력하기 때문이죠. 어색하지 않을까 걱정했지만, 이제는 업계의 표준이 되었죠. 많은 앱에서 이 형태를 적용하고 있어요. 처음 이 화면을 디자인하게 된 과정을 소개해드릴게요.
Link:  https://toss.tech/article/toss-signup-process
Published:  Tue, 20 Sep 2022 04:55:25 GMT

GUID:  https://toss.tech/article/introducing-toss-error-message-system
Title:  가이드라인을 시스템으로 만드는 법
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-6-og.png?v=2
Description :  UX 라이팅은 한 명이 아닌, 팀원 모두가 잘 쓰는 게 중요해요. 이전까지는 세션이나 가이드라인 등의 방법으로 이 문제를 접근해왔지만, 토스는 조금 다르게 풀어봤어요. 시스템으로요.
Link:  https://toss.tech/article/introducing-toss-error-message-system
Published:  Fri, 04 Nov 2022 07:39:40 GMT

GUID:  https://toss.tech/article/signup
Title:  완성 없는 이야기, 가입 과정 개선
Thumbnail:  https://static.toss.im/simplicity23/graphics/sim23-signup-og.png
Description :  가입 완료율을 높이기 위해 고민하고 시도했던 것들을 이야기 해요.
Link:  https://toss.tech/article/signup
Published:  Tue, 12 Sep 2023 13:44:00 GMT

GUID:  https://toss.tech/article/insurance-claim-process
Title:  토스 디자인 원칙, Easy to answer
Thumbnail:  https://toss.tech/wp-content/uploads/2022/12/Columns.png
Description :  토스에는 좋은 제품을 만들기 위한 디자인 원칙들이 있어요. 오늘은 그 중에 easy to answer라는 항목을 적용해서 제품을 개선한 사례를 소개해드릴게요.
Link:  https://toss.tech/article/insurance-claim-process
Published:  Wed, 14 Dec 2022 12:45:20 GMT

GUID:  https://toss.tech/article/how-to-make-marvelous-productivity-with-well-made-system
Title:  잘 만든 시스템 하나로 미친 효율 얻는 방법
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/design-article-13.png?v=2
Description :  ‘원래 하던 일’을 단순히 똑같이 해내는 것이 아니라, ‘더 잘하게 만드는 도구’를 만들어서 원래 하던 일을 훨씬 더 빠른 속도로 해낼 수 있어요. 효율이 복리로 쌓이게 만드는 거죠.
Link:  https://toss.tech/article/how-to-make-marvelous-productivity-with-well-made-system
Published:  Thu, 12 Jan 2023 10:05:56 GMT

GUID:  https://toss.tech/article/1st-product-designer-tools
Title:  토스 최초의 Product Designer(Tools)의 일하는 방식
Thumbnail:  https://static.toss.im/illusts/1st-product-designer-tools.png
Description :  토스팀에서 첫 Product Designer (Tools) 직무로 일하기 시작하면서 어떤 걸 배웠는지 알려드릴게요.
Link:  https://toss.tech/article/1st-product-designer-tools
Published:  Tue, 20 Jun 2023 07:35:00 GMT

GUID:  https://toss.tech/article/tosssec-qateam
Title:  토스증권 QA Team을 소개합니다
Thumbnail:  https://og.toss.tech?title=%ED%86%A0%EC%8A%A4%EC%A6%9D%EA%B6%8C%20QA%20Team%EC%9D%84%20%EC%86%8C%EA%B0%9C%ED%95%A9%EB%8B%88%EB%8B%A4&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Ftoss-tech%2Fqa-og.png&v=2
Description :  이 글은 토스증권 QA Team에 입사를 고려 중인 지원자들을 위해 작성된 글입니다. 토스증권 QA Manager 하는 역할과 일하는 방식은 어떻게 다를까요?
Link:  https://toss.tech/article/tosssec-qateam
Published:  Tue, 25 Oct 2022 06:48:53 GMT

GUID:  https://toss.tech/article/frontend-diving-club
Title:  놀러오세요! 프론트엔드 다이빙 클럽
Thumbnail:  https://static.toss.im/assets/toss-tech/frontend-diving-club.jpg
Description :  프론트엔드에 관한 깊은 이야기를 나눌 수 있는 오프라인 커뮤니티, 프론트엔드 다이빙 클럽을 소개합니다.
Link:  https://toss.tech/article/frontend-diving-club
Published:  Fri, 21 Jul 2023 04:47:00 GMT

GUID:  https://toss.tech/article/toss-frontend-chapter
Title:  토스 프론트엔드 챕터를 소개합니다!
Thumbnail:  https://og.toss.tech?title=%ED%86%A0%EC%8A%A4%20%ED%94%84%EB%A1%A0%ED%8A%B8%EC%97%94%EB%93%9C%20%EB%B8%94%EB%A1%9C%EA%B7%B8%EC%97%90%0D%0A%EC%98%A4%EC%8B%A0%20%EA%B2%83%EC%9D%84%20%ED%99%98%EC%98%81%ED%95%A9%EB%8B%88%EB%8B%A4%21&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Ftech-blog%2Fog-image%3A%2Ftechblog-01-about-fechapter-og.png&v=2
Description :  토스에서 프론트엔드 개발자가 일하는 방법과 맡고 있는 역할에 대해 소개드립니다.
Link:  https://toss.tech/article/toss-frontend-chapter
Published:  Tue, 27 Apr 2021 23:00:00 GMT

GUID:  https://toss.tech/article/1st-graphic-designer
Title:  첫 그래픽 디자이너가 했던 가장 비효율적인 일
Thumbnail:  https://static.toss.im/assets/toss-tech/gd_og.png
Description :  첫 번째 그래픽 디자이너로 입사해서 자리잡기까지의 여정과, 그래픽 디자이너만의 노하우를 알려드려요.
Link:  https://toss.tech/article/1st-graphic-designer
Published:  Thu, 20 Apr 2023 03:00:00 GMT

GUID:  https://toss.tech/article/real-reason
Title:  문제 원인의 원인을 찾아서
Thumbnail:  https://static.toss.im/assets/homepage/tosstech/og/Frame%2045.png?v=2
Description :  좋은 해결책을 내기 위해서 제가 쓰는 방법은, 문제 원인의 원인을 찾는 거예요. 진짜 문제를 발견하면, 임팩트 있는 해결책을 생각해 낼 확률이 훨씬 높아지죠.
Link:  https://toss.tech/article/real-reason
Published:  Thu, 02 Mar 2023 08:40:06 GMT

GUID:  https://toss.tech/article/dev-agility
Title:  개발자의 애질리티
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-06-agility-og.png?v=2
Description :  이 글은 토스페이먼츠에 입사하신, 혹은 입사를 고려 중인 개발자분들을 위해 작성된 글입니다. 애자일하게 일한다는 것은 어떠한 의미일까요?
Link:  https://toss.tech/article/dev-agility
Published:  Thu, 14 Oct 2021 13:32:16 GMT

GUID:  https://toss.tech/article/22439
Title:  슬랙봇 디자인 101
Thumbnail:  https://static.toss.im/career-resource/slackbot101_og.png
Description :  슬랙봇 디자인을 잘 하려면 무엇을 알아야하는지 알려드릴게요.
Link:  https://toss.tech/article/22439
Published:  Thu, 17 Aug 2023 08:38:00 GMT

GUID:  https://toss.tech/article/typescript-type-compatibility
Title:  TypeScript 타입 시스템 뜯어보기: 타입 호환성
Thumbnail:  https://og.toss.tech?title=TypeScript%0D%0A%ED%83%80%EC%9E%85%20%EC%8B%9C%EC%8A%A4%ED%85%9C%20%EB%9C%AF%EC%96%B4%EB%B3%B4%EA%B8%B0%3A%0D%0A%ED%83%80%EC%9E%85%20%ED%98%B8%ED%99%98%EC%84%B1&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Ftoss-tech%2Ftech-article-ts-og.png&v=2
Description :  타입호환성은 무엇이며 왜 필요할까요? 타입호환이 지원되지 않는 경우가 존재한다는 것을 아셨나요? 평소 익숙했던 개념들에 대해 질문을 던져가며 TypeScript 타입 시스템에 관해 심도있게 알아보고자 합니다.
Link:  https://toss.tech/article/typescript-type-compatibility
Published:  Tue, 25 Oct 2022 15:38:54 GMT

GUID:  https://toss.tech/article/slash23-security
Title:  금융사 최초의 Zero Trust 아키텍처 도입기
Thumbnail:  https://static.toss.im/career-resource/techblog_slash23_og_03%20(1).png
Description :  왜 보안이 강화될수록, 업무는 불편해지는 걸까요? 
금융에서 가장 중요한 가치인 ‘안전’과 혁신을 위한 ‘업무 편의성’, 두 마리 토끼를 모두 잡기 위한 여정을 소개합니다.  
Link:  https://toss.tech/article/slash23-security
Published:  Fri, 01 Sep 2023 02:59:00 GMT

GUID:  https://toss.tech/article/slash23-devops
Title:  유연하고 안전하게 배포 Pipeline 운영하기
Thumbnail:  https://static.toss.im/career-resource/techblog_slash23_og_06_김동석.png
Description :  생산성, 안정성, 보안성을 모두 달성하기 위해선 ‘Pipeline’을 어떻게 설정해야 할까요?
다양하고 복잡한 토스뱅크의 배포 ‘Pipeline’를 안전하게 운영하기 위한 노력을 소개합니다.
Link:  https://toss.tech/article/slash23-devops
Published:  Thu, 12 Oct 2023 11:30:00 GMT

GUID:  https://toss.tech/article/mydoc
Title:  크고 복잡한 제품, 과감하게 갈아엎기
Thumbnail:  https://static.toss.im/career-resource/sim23-mydoc-og.png
Description :  디자이너들도 어떻게 하면 빠른 속도를 내며 전체적인 경험을 맞출 수 있을지 매 순간 고민하고 있어요. 
Link:  https://toss.tech/article/mydoc
Published:  Sun, 03 Sep 2023 13:58:00 GMT

GUID:  https://toss.tech/article/node-modules-and-yarn-berry
Title:  node_modules로부터 우리를 구원해 줄 Yarn Berry
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-04-yarn-berry-og.png?v=2
Description :  토스 프론트엔드 레포지토리 대부분에서 사용하고 있는 패키지 매니저 Yarn Berry. 채택하게 된 배경과 사용하면서 좋았던 점을 공유합니다.
Link:  https://toss.tech/article/node-modules-and-yarn-berry
Published:  Fri, 07 May 2021 08:18:00 GMT

GUID:  https://toss.tech/article/jscodeshift
Title:  JSCodeShift로 기술 부채 청산하기
Thumbnail:  https://static.toss.im/assets/tech-blog/og-image:/techblog-03-jscodeshift-og.png?v=2
Description :  기술 부채는 개발할수록 쌓여만 갑니다. 프론트엔드 챕터가 JSCodeShift를 이용하여 순식간에 100개 서비스의 기술 부채를 해결한 경험을 소개합니다.
Link:  https://toss.tech/article/jscodeshift
Published:  Tue, 04 May 2021 08:16:00 GMT

GUID:  https://toss.tech/article/commonjs-esm-exports-field
Title:  CommonJS와 ESM에 모두 대응하는 라이브러리 개발하기: exports field
Thumbnail:  https://og.toss.tech?title=CommonJS%EC%99%80%20ESM%EC%97%90%20%EB%AA%A8%EB%91%90%20%EB%8C%80%EC%9D%91%ED%95%98%EB%8A%94%20%EB%9D%BC%EC%9D%B4%EB%B8%8C%EB%9F%AC%EB%A6%AC%20%EA%B0%9C%EB%B0%9C%ED%95%98%EA%B8%B0%3A%20exports%20field&imageUrl=https%3A%2F%2Fstatic.toss.im%2Fassets%2Fhomepage%2Ftosstech%2Fog%2Ftechblog-11-node-js-og.png&v=2
Description :  Node.js에는 두 가지 Module System이 존재합니다. 토스 프론트엔드 챕터에서 운영하는 100개가 넘는 라이브러리들은 그것에 어떻게 대응하고 있을까요?
Link:  https://toss.tech/article/commonjs-esm-exports-field
Published:  Tue, 04 Oct 2022 08:37:32 GMT
```

# Execution Time
```
Two attempts

Total: 64 Feeds

Execution Time: 4.326115625s
Execution Time: 4.204783958s

The time it took per feed is 0.06664765299s
```
