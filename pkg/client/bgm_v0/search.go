package bgm_v0

import (
	"context"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/errno"
	"strconv"
)

//SearchSmallSubjectByKeywords Bangumi OpenAPI 搜索 return small type
func (cli Client) SearchSmallSubjectByKeywords(ctx context.Context, authToken string, keywords string, subjectType items.SubjectType,
	responseGroup string, start int64, maxResults int64) ([]*items.SubjectSmall, error) {

	if keywords == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "keywords is required!")
	}
	if subjectType != 0 && !subjectType.IsSupported() {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject type is invalid!")
	}
	if start < 0 {
		return nil, errno.Errorf(errno.ErrBadRequest, "start is invalid!")
	}
	if maxResults <= 0 {
		return nil, errno.Errorf(errno.ErrBadRequest, "maxResults is invalid!")
	} else if maxResults > 25 {
		maxResults = 25
	}

	resp := struct {
		Result int
		List   []*items.SubjectSmall
	}{}

	param := map[string]string{
		"type":          subjectType.ToString(),
		"responseGroup": "small",
		"start":         strconv.FormatInt(start, 10),
		"max_results":   strconv.FormatInt(maxResults, 10),
	}
	err := cli.GET(ctx, "/search/subject/"+keywords, authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.List, nil
}

//SearchMediumSubjectByKeywords Bangumi OpenAPI 搜索 return medium type
func (cli Client) SearchMediumSubjectByKeywords(ctx context.Context, authToken string, keywords string, subjectType items.SubjectType,
	responseGroup string, start int64, maxResults int64) ([]*items.SubjectMedium, error) {
	if keywords == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "keywords is required!")
	}
	if subjectType != 0 && !subjectType.IsSupported() {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject type is invalid!")
	}
	if start < 0 {
		return nil, errno.Errorf(errno.ErrBadRequest, "start is invalid!")
	}
	if maxResults < 0 {
		return nil, errno.Errorf(errno.ErrBadRequest, "maxResults is invalid!")
	} else if maxResults > 25 || maxResults == 0 {
		maxResults = 25
	}

	resp := struct {
		Result int
		List   []*items.SubjectMedium
	}{}
	param := map[string]string{
		"type":          subjectType.ToString(),
		"responseGroup": "small",
		"start":         strconv.FormatInt(start, 10),
		"max_results":   strconv.FormatInt(maxResults, 10),
	}
	err := cli.GET(ctx, "/search/subject/"+keywords, authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.List, nil
}
