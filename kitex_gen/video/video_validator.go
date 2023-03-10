// Code generated by Validator v0.1.4. DO NOT EDIT.

package video

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *Video) IsValid() error {
	if p.Author != nil {
		if err := p.Author.IsValid(); err != nil {
			return fmt.Errorf("filed Author not valid, %w", err)
		}
	}
	return nil
}
func (p *User) IsValid() error {
	return nil
}
func (p *Comment) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
func (p *FeedRequest) IsValid() error {
	if p.Token != nil {
		if err := p.Token.IsValid(); err != nil {
			return fmt.Errorf("filed Token not valid, %w", err)
		}
	}
	return nil
}
func (p *FeedResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *PublishRequest) IsValid() error {
	if p.Token != nil {
		if err := p.Token.IsValid(); err != nil {
			return fmt.Errorf("filed Token not valid, %w", err)
		}
	}
	if len(p.Data) < int(1) {
		return fmt.Errorf("field Data min_len rule failed, current value: %d", len(p.Data))
	}
	if len(p.Title) < int(1) {
		return fmt.Errorf("field Title min_len rule failed, current value: %d", len(p.Title))
	}
	return nil
}
func (p *PublishResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *PublishListRequest) IsValid() error {
	if p.Token != nil {
		if err := p.Token.IsValid(); err != nil {
			return fmt.Errorf("filed Token not valid, %w", err)
		}
	}
	return nil
}
func (p *PublishListResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentActionRequest) IsValid() error {
	if p.Token != nil {
		if err := p.Token.IsValid(); err != nil {
			return fmt.Errorf("filed Token not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentActionResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	if p.Comment != nil {
		if err := p.Comment.IsValid(); err != nil {
			return fmt.Errorf("filed Comment not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentListRequest) IsValid() error {
	if p.Token != nil {
		if err := p.Token.IsValid(); err != nil {
			return fmt.Errorf("filed Token not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentListResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *FavoriteActionRequest) IsValid() error {
	if p.Token != nil {
		if err := p.Token.IsValid(); err != nil {
			return fmt.Errorf("filed Token not valid, %w", err)
		}
	}
	return nil
}
func (p *FavoriteActionResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *FavoriteListRequest) IsValid() error {
	if p.Token != nil {
		if err := p.Token.IsValid(); err != nil {
			return fmt.Errorf("filed Token not valid, %w", err)
		}
	}
	return nil
}
func (p *FavoriteListResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
