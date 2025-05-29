// Package wemeetopenapi is auto generate
/*
    腾讯会议OpenAPI

    SAAS版RESTFUL风格API

    API version: v1.0.8
*/
package wemeetopenapi


// V1SmartChaptersGet200Response struct for V1SmartChaptersGet200Response
type V1SmartChaptersGet200Response struct {
    // ChapterList对象数组
    ChapterList []V1SmartChaptersGet200ResponseChapterListInner `json:"chapter_list,omitempty"`
}

// V1SmartChaptersGet200ResponseChapterListInner struct for V1SmartChaptersGet200ResponseChapterListInner
type V1SmartChaptersGet200ResponseChapterListInner struct {
    // 章节唯一ID
    ChapterId *string `json:"chapter_id,omitempty"`
    // 章节主题，返回base64编码后的结果
    ChapterName *string `json:"chapter_name,omitempty"`
    // 章节封面图片url
    PicUrl *string `json:"pic_url,omitempty"`
    // 开始时间戳（单位毫秒）
    StartTime *string `json:"start_time,omitempty"`
}

// V1SmartFullsummaryGet200Response struct for V1SmartFullsummaryGet200Response
type V1SmartFullsummaryGet200Response struct {
    // 智能总结内容
    AiSummary *string `json:"ai_summary,omitempty"`
}

// V1SmartSpeakersGet200Response struct for V1SmartSpeakersGet200Response
type V1SmartSpeakersGet200Response struct {
    // 本录制文件的发言人列表，以对象数组形式返回
    SpeakerList []V1SmartSpeakersGet200ResponseSpeakerListInner `json:"speaker_list,omitempty"`
    // 发言人总数
    TotalCount *int64 `json:"total_count,omitempty"`
}

// V1SmartSpeakersGet200ResponseSpeakerListInner struct for V1SmartSpeakersGet200ResponseSpeakerListInner
type V1SmartSpeakersGet200ResponseSpeakerListInner struct {
    // 会议中为每个参会成员授予的临时 ID，以会议为维度，表示同一场会议内用户的唯一标识，不同会议间 ms_open_id 隔离。
    MsOpenId *string `json:"ms_open_id,omitempty"`
    // 发言人ID。speaker_id 必须与 speaker_id_type 配合使用。根据 speaker_id_type 的值，speaker_id 代表不同类型。
    SpeakerId *string `json:"speaker_id,omitempty"`
    // 发言人ID类型： 1：userid 2：openid 6：temp_id（临时 ID，上传的文件无法映射到 userid，故仅在当前录制发言人中代表唯一标识）
    SpeakerIdType *int64 `json:"speaker_id_type,omitempty"`
    // 发言人名称，base64编码
    SpeakerName *string `json:"speaker_name,omitempty"`
    // 本录制文件某个具体发言人的发言时间段，以对象数组形式返回
    SpeakerTime []V1SmartSpeakersGet200ResponseSpeakerListInnerSpeakerTimeInner `json:"speaker_time,omitempty"`
    // 发言总时长
    TotalTime *int64 `json:"total_time,omitempty"`
}

// V1SmartSpeakersGet200ResponseSpeakerListInnerSpeakerTimeInner struct for V1SmartSpeakersGet200ResponseSpeakerListInnerSpeakerTimeInner
type V1SmartSpeakersGet200ResponseSpeakerListInnerSpeakerTimeInner struct {
    // 结束时间戳（单位毫秒）
    EndTime *string `json:"end_time,omitempty"`
    // 发言语句id
    Sid *string `json:"sid,omitempty"`
    // 开始时间戳（单位毫秒）
    StartTime *string `json:"start_time,omitempty"`
}

// V1SmartTopicsGet200Response struct for V1SmartTopicsGet200Response
type V1SmartTopicsGet200Response struct {
    // 本录制文件的智能话题列表，以对象数组形式返回
    AiTopicList []V1SmartTopicsGet200ResponseAiTopicListInner `json:"ai_topic_list,omitempty"`
}

// V1SmartTopicsGet200ResponseAiTopicListInner struct for V1SmartTopicsGet200ResponseAiTopicListInner
type V1SmartTopicsGet200ResponseAiTopicListInner struct {
    // 话题唯一ID
    TopicId *string `json:"topic_id,omitempty"`
    // 话题主题，base6编码
    TopicName *string `json:"topic_name,omitempty"`
    // 本话题的发言段落及时间段，以对象数组形式返回
    TopicTime []V1SmartTopicsGet200ResponseAiTopicListInnerTopicTimeInner `json:"topic_time,omitempty"`
}

// V1SmartTopicsGet200ResponseAiTopicListInnerTopicTimeInner struct for V1SmartTopicsGet200ResponseAiTopicListInnerTopicTimeInner
type V1SmartTopicsGet200ResponseAiTopicListInnerTopicTimeInner struct {
    // 结束时间戳（单位毫秒）
    EndTime *string `json:"end_time,omitempty"`
    // 段落ID
    Pid *string `json:"pid,omitempty"`
    // 开始时间戳（单位毫秒）
    StartTime *string `json:"start_time,omitempty"`
}
