package out_test

import (
	"errors"
	"fmt"

	"github.com/Medium/medium-sdk-go"
	"github.com/cappyzawa/medium-resource"
	"github.com/cappyzawa/medium-resource/fakes"
	"github.com/cappyzawa/medium-resource/out"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Command", func() {
	var (
		fakeMC    *fakes.FakeMediumClient
		command   out.Command
		sourceDir string
		request   out.Request
	)

	BeforeEach(func() {
		fakeMC = new(fakes.FakeMediumClient)
		command = out.Command{
			MediumClient: fakeMC,
		}
		sourceDir = "../testdata"
		request = out.Request{
			Source: resource.Source{
				AccessToken: "token",
			},
			Params: out.Params{
				Format:       "markdown",
				ContentFile:  "out/content_file.md",
				Tags:         []string{"tag1", "tag2"},
				CanonicalURL: "https://canonical.com",
				Status:       "draft",
				License:      "mit",
			},
		}
	})

	Describe("Run()", func() {
		BeforeEach(func() {
			fakeMC.GetUserReturns(&medium.User{ID: "id"}, nil)
			fakeMC.CreatePostReturns(&medium.Post{
				ID:           "id",
				Title:        "title",
				PublishState: "draft",
				URL:          "https://posted.com",
			}, nil)
		})
		Context("when token is empty", func() {
			BeforeEach(func() {
				request.Source.AccessToken = ""
			})
			It("error should occur", func() {
				res, err := command.Run(sourceDir, request)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})
		})
		Context("when content_file is empty", func() {
			BeforeEach(func() {
				request.Params.ContentFile = ""
			})
			It("error should occur", func() {
				res, err := command.Run(sourceDir, request)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})
		})
		Context("when token is invalid", func() {
			It("error should occur", func() {
				fakeMC.GetUserReturns(nil, errors.New("unauthorized"))
				res, err := command.Run(sourceDir, request)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})
		})
		Context("when title is empty", func() {
			BeforeEach(func() {
				request.Params.Title = ""
			})
			It("the first line of content_file is used", func() {
				fakeMC.GetUserReturns(&medium.User{ID: "id"}, nil)
				_, err := command.Run(sourceDir, request)
				Expect(err).NotTo(HaveOccurred())
				option := fakeMC.CreatePostArgsForCall(0)
				Expect(option.Title).To(Equal("item1"))
			})
		})
		Context("when format is empty", func() {
			BeforeEach(func() {
				request.Params.Format = ""
			})
			It("default format is markdown", func() {
				fakeMC.GetUserReturns(&medium.User{ID: "id"}, nil)
				_, err := command.Run(sourceDir, request)
				Expect(err).NotTo(HaveOccurred())
				option := fakeMC.CreatePostArgsForCall(0)
				Expect(option.ContentFormat).To(Equal(medium.ContentFormat(medium.ContentFormatMarkdown)))
			})
		})
		Context("when status is empty", func() {
			BeforeEach(func() {
				request.Params.Status = ""
			})
			It("default status is draft", func() {
				fakeMC.GetUserReturns(&medium.User{ID: "id"}, nil)
				_, err := command.Run(sourceDir, request)
				Expect(err).NotTo(HaveOccurred())
				option := fakeMC.CreatePostArgsForCall(0)
				Expect(option.PublishStatus).To(Equal(medium.PublishStatus(medium.PublishStatusDraft)))
			})
		})
		It("content_file is used as content", func() {
			fakeMC.GetUserReturns(&medium.User{ID: "id"}, nil)
			_, err := command.Run(sourceDir, request)
			Expect(err).NotTo(HaveOccurred())
			option := fakeMC.CreatePostArgsForCall(0)
			Expect(option.Content).To(Equal("## item2\ncontent\n"))
		})
		It("response contains version & metadata", func() {
			res, err := command.Run(sourceDir, request)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Version.ID).To(Equal("id"))
			Expect(len(res.Metadata)).NotTo(BeZero())
			Expect(res.Metadata[0].Name).To(Equal("title"))
			Expect(res.Metadata[0].Value).To(Equal("title"))
			Expect(res.Metadata[1].Name).To(Equal("status"))
			Expect(res.Metadata[1].Value).To(Equal("draft"))
			Expect(res.Metadata[2].Name).To(Equal("url"))
			Expect(res.Metadata[2].Value).To(Equal("https://posted.com"))
		})
	})

	Describe("ExtractTitleAndContent()", func() {
		var path string
		Context("file exists", func() {
			BeforeEach(func() {
				path = fmt.Sprintf("%s/%s", sourceDir, request.Params.ContentFile)
			})
			It("title is first lint of file", func() {
				title, content, err := command.ExtractTitleAndContent(path)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(title)).To(Equal("item1"))
				Expect(string(content)).To(Equal("## item2\ncontent\n"))
			})
		})
		Context("file does not exist", func() {
			BeforeEach(func() {
				path = "missing/missing"
			})
			It("error should occur", func() {
				title, content, err := command.ExtractTitleAndContent(path)
				Expect(err).To(HaveOccurred())
				Expect(title).To(BeEmpty())
				Expect(content).To(BeEmpty())
			})
		})
	})

})
