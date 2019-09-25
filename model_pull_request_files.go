/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// Pull Request Commit文件列表。最多显示300条diff
type PullRequestFiles struct {
	Sha string `json:"sha,omitempty"`
	Filename string `json:"filename,omitempty"`
	Status string `json:"status,omitempty"`
	Additions string `json:"additions,omitempty"`
	Deletions string `json:"deletions,omitempty"`
	BlobUrl string `json:"blob_url,omitempty"`
	RawUrl string `json:"raw_url,omitempty"`
	Patch string `json:"patch,omitempty"`
}