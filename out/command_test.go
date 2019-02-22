package out_test

import (
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
		Context("when title is empty", func() {
			BeforeEach(func() {
				request.Params.Title = ""
			})
			It("the first line of content_file is used", func() {
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
				_, err := command.Run(sourceDir, request)
				Expect(err).NotTo(HaveOccurred())
				option := fakeMC.CreatePostArgsForCall(0)
				Expect(option.ContentFormat).To(Equal(medium.ContentFormatMarkdown))
			})
		})
		Context("when status is empty", func() {
			BeforeEach(func() {
				request.Params.Status = ""
			})
			It("default status is draft", func() {
				_, err := command.Run(sourceDir, request)
				Expect(err).NotTo(HaveOccurred())
				option := fakeMC.CreatePostArgsForCall(0)
				Expect(option.PublishStatus).To(Equal(medium.PublishStatusDraft))
			})
		})
		It("content_file is used as content", func() {
			_, err := command.Run(sourceDir, request)
			Expect(err).NotTo(HaveOccurred())
			option := fakeMC.CreatePostArgsForCall(0)
			Expect(option.Content).To(Equal("##item2\ncontent"))
		})
	})

})
