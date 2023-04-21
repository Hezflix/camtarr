package tautulli

type respGetActivity struct {
	Response struct {
		Result  string `json:"result"`
		Message any    `json:"message"`
		Data    struct {
			StreamCount string `json:"stream_count"`
			Sessions    []struct {
				SessionKey              string   `json:"session_key"`
				MediaType               string   `json:"media_type"`
				ViewOffset              string   `json:"view_offset"`
				ProgressPercent         string   `json:"progress_percent"`
				QualityProfile          string   `json:"quality_profile"`
				SyncedVersionProfile    string   `json:"synced_version_profile"`
				OptimizedVersionProfile string   `json:"optimized_version_profile"`
				User                    string   `json:"user"`
				ChannelStream           int      `json:"channel_stream"`
				SectionID               string   `json:"section_id"`
				LibraryName             string   `json:"library_name"`
				RatingKey               string   `json:"rating_key"`
				ParentRatingKey         string   `json:"parent_rating_key"`
				GrandparentRatingKey    string   `json:"grandparent_rating_key"`
				Title                   string   `json:"title"`
				ParentTitle             string   `json:"parent_title"`
				GrandparentTitle        string   `json:"grandparent_title"`
				OriginalTitle           string   `json:"original_title"`
				SortTitle               string   `json:"sort_title"`
				EditionTitle            string   `json:"edition_title"`
				MediaIndex              string   `json:"media_index"`
				ParentMediaIndex        string   `json:"parent_media_index"`
				Studio                  string   `json:"studio"`
				ContentRating           string   `json:"content_rating"`
				Summary                 string   `json:"summary"`
				Tagline                 string   `json:"tagline"`
				Rating                  string   `json:"rating"`
				RatingImage             string   `json:"rating_image"`
				AudienceRating          string   `json:"audience_rating"`
				AudienceRatingImage     string   `json:"audience_rating_image"`
				UserRating              string   `json:"user_rating"`
				Duration                string   `json:"duration"`
				Year                    string   `json:"year"`
				ParentYear              string   `json:"parent_year"`
				GrandparentYear         string   `json:"grandparent_year"`
				Thumb                   string   `json:"thumb"`
				ParentThumb             string   `json:"parent_thumb"`
				GrandparentThumb        string   `json:"grandparent_thumb"`
				Art                     string   `json:"art"`
				Banner                  string   `json:"banner"`
				OriginallyAvailableAt   string   `json:"originally_available_at"`
				AddedAt                 string   `json:"added_at"`
				UpdatedAt               string   `json:"updated_at"`
				LastViewedAt            string   `json:"last_viewed_at"`
				GUID                    string   `json:"guid"`
				ParentGUID              string   `json:"parent_guid"`
				GrandparentGUID         string   `json:"grandparent_guid"`
				Directors               []string `json:"directors"`
				Writers                 []string `json:"writers"`
				Actors                  []string `json:"actors"`
				Genres                  []string `json:"genres"`
				Labels                  []any    `json:"labels"`
				Collections             []any    `json:"collections"`
				Guids                   []string `json:"guids"`
				Markers                 []struct {
					ID              int    `json:"id"`
					Type            string `json:"type"`
					StartTimeOffset int    `json:"start_time_offset"`
					EndTimeOffset   int    `json:"end_time_offset"`
					First           bool   `json:"first"`
					Final           bool   `json:"final"`
				} `json:"markers"`
				ParentGuids                  []any    `json:"parent_guids"`
				GrandparentGuids             []any    `json:"grandparent_guids"`
				FullTitle                    string   `json:"full_title"`
				ChildrenCount                int      `json:"children_count"`
				Live                         int      `json:"live"`
				ID                           string   `json:"id"`
				Container                    string   `json:"container"`
				Bitrate                      string   `json:"bitrate"`
				Height                       string   `json:"height"`
				Width                        string   `json:"width"`
				AspectRatio                  string   `json:"aspect_ratio"`
				VideoCodec                   string   `json:"video_codec"`
				VideoResolution              string   `json:"video_resolution"`
				VideoFullResolution          string   `json:"video_full_resolution"`
				VideoFramerate               string   `json:"video_framerate"`
				VideoProfile                 string   `json:"video_profile"`
				AudioCodec                   string   `json:"audio_codec"`
				AudioChannels                string   `json:"audio_channels"`
				AudioChannelLayout           string   `json:"audio_channel_layout"`
				AudioProfile                 string   `json:"audio_profile"`
				OptimizedVersion             int      `json:"optimized_version"`
				ChannelCallSign              string   `json:"channel_call_sign"`
				ChannelIdentifier            string   `json:"channel_identifier"`
				ChannelThumb                 string   `json:"channel_thumb"`
				File                         string   `json:"file"`
				FileSize                     string   `json:"file_size"`
				Indexes                      int      `json:"indexes"`
				Selected                     int      `json:"selected"`
				Type                         string   `json:"type"`
				VideoCodecLevel              string   `json:"video_codec_level"`
				VideoBitrate                 string   `json:"video_bitrate"`
				VideoBitDepth                string   `json:"video_bit_depth"`
				VideoChromaSubsampling       string   `json:"video_chroma_subsampling"`
				VideoColorPrimaries          string   `json:"video_color_primaries"`
				VideoColorRange              string   `json:"video_color_range"`
				VideoColorSpace              string   `json:"video_color_space"`
				VideoColorTrc                string   `json:"video_color_trc"`
				VideoDynamicRange            string   `json:"video_dynamic_range"`
				VideoFrameRate               string   `json:"video_frame_rate"`
				VideoRefFrames               string   `json:"video_ref_frames"`
				VideoHeight                  string   `json:"video_height"`
				VideoWidth                   string   `json:"video_width"`
				VideoLanguage                string   `json:"video_language"`
				VideoLanguageCode            string   `json:"video_language_code"`
				VideoScanType                string   `json:"video_scan_type"`
				AudioBitrate                 string   `json:"audio_bitrate"`
				AudioBitrateMode             string   `json:"audio_bitrate_mode"`
				AudioSampleRate              string   `json:"audio_sample_rate"`
				AudioLanguage                string   `json:"audio_language"`
				AudioLanguageCode            string   `json:"audio_language_code"`
				SubtitleCodec                string   `json:"subtitle_codec"`
				SubtitleContainer            string   `json:"subtitle_container"`
				SubtitleFormat               string   `json:"subtitle_format"`
				SubtitleForced               int      `json:"subtitle_forced"`
				SubtitleLocation             string   `json:"subtitle_location"`
				SubtitleLanguage             string   `json:"subtitle_language"`
				SubtitleLanguageCode         string   `json:"subtitle_language_code"`
				RowID                        int      `json:"row_id"`
				UserID                       int      `json:"user_id"`
				Username                     string   `json:"username"`
				FriendlyName                 string   `json:"friendly_name"`
				UserThumb                    string   `json:"user_thumb"`
				Email                        string   `json:"email"`
				IsActive                     int      `json:"is_active"`
				IsAdmin                      int      `json:"is_admin"`
				IsHomeUser                   int      `json:"is_home_user"`
				IsAllowSync                  int      `json:"is_allow_sync"`
				IsRestricted                 int      `json:"is_restricted"`
				DoNotify                     int      `json:"do_notify"`
				KeepHistory                  int      `json:"keep_history"`
				DeletedUser                  int      `json:"deleted_user"`
				AllowGuest                   int      `json:"allow_guest"`
				SharedLibraries              []string `json:"shared_libraries"`
				LastSeen                     any      `json:"last_seen"`
				IPAddress                    string   `json:"ip_address"`
				IPAddressPublic              string   `json:"ip_address_public"`
				Device                       string   `json:"device"`
				Platform                     string   `json:"platform"`
				PlatformName                 string   `json:"platform_name"`
				PlatformVersion              string   `json:"platform_version"`
				Product                      string   `json:"product"`
				ProductVersion               string   `json:"product_version"`
				Profile                      string   `json:"profile"`
				Player                       string   `json:"player"`
				MachineID                    string   `json:"machine_id"`
				State                        string   `json:"state"`
				Local                        int      `json:"local"`
				Relayed                      int      `json:"relayed"`
				Secure                       int      `json:"secure"`
				SessionID                    string   `json:"session_id"`
				Bandwidth                    string   `json:"bandwidth"`
				Location                     string   `json:"location"`
				TranscodeKey                 string   `json:"transcode_key"`
				TranscodeThrottled           int      `json:"transcode_throttled"`
				TranscodeProgress            int      `json:"transcode_progress"`
				TranscodeSpeed               string   `json:"transcode_speed"`
				TranscodeAudioChannels       string   `json:"transcode_audio_channels"`
				TranscodeAudioCodec          string   `json:"transcode_audio_codec"`
				TranscodeVideoCodec          string   `json:"transcode_video_codec"`
				TranscodeWidth               string   `json:"transcode_width"`
				TranscodeHeight              string   `json:"transcode_height"`
				TranscodeContainer           string   `json:"transcode_container"`
				TranscodeProtocol            string   `json:"transcode_protocol"`
				TranscodeMinOffsetAvailable  int      `json:"transcode_min_offset_available"`
				TranscodeMaxOffsetAvailable  int      `json:"transcode_max_offset_available"`
				TranscodeHwRequested         int      `json:"transcode_hw_requested"`
				TranscodeHwDecode            string   `json:"transcode_hw_decode"`
				TranscodeHwDecodeTitle       string   `json:"transcode_hw_decode_title"`
				TranscodeHwEncode            string   `json:"transcode_hw_encode"`
				TranscodeHwEncodeTitle       string   `json:"transcode_hw_encode_title"`
				TranscodeHwFullPipeline      int      `json:"transcode_hw_full_pipeline"`
				AudioDecision                string   `json:"audio_decision"`
				VideoDecision                string   `json:"video_decision"`
				SubtitleDecision             string   `json:"subtitle_decision"`
				Throttled                    string   `json:"throttled"`
				TranscodeHwDecoding          int      `json:"transcode_hw_decoding"`
				TranscodeHwEncoding          int      `json:"transcode_hw_encoding"`
				StreamContainer              string   `json:"stream_container"`
				StreamBitrate                string   `json:"stream_bitrate"`
				StreamAspectRatio            string   `json:"stream_aspect_ratio"`
				StreamVideoFramerate         string   `json:"stream_video_framerate"`
				StreamVideoResolution        string   `json:"stream_video_resolution"`
				StreamDuration               string   `json:"stream_duration"`
				StreamContainerDecision      string   `json:"stream_container_decision"`
				OptimizedVersionTitle        string   `json:"optimized_version_title"`
				SyncedVersion                int      `json:"synced_version"`
				LiveUUID                     string   `json:"live_uuid"`
				BifThumb                     string   `json:"bif_thumb"`
				Subtitles                    int      `json:"subtitles"`
				TranscodeDecision            string   `json:"transcode_decision"`
				ContainerDecision            string   `json:"container_decision"`
				StreamVideoFullResolution    string   `json:"stream_video_full_resolution"`
				StreamVideoBitrate           string   `json:"stream_video_bitrate"`
				StreamVideoBitDepth          string   `json:"stream_video_bit_depth"`
				StreamVideoChromaSubsampling string   `json:"stream_video_chroma_subsampling"`
				StreamVideoCodec             string   `json:"stream_video_codec"`
				StreamVideoCodecLevel        string   `json:"stream_video_codec_level"`
				StreamVideoColorPrimaries    string   `json:"stream_video_color_primaries"`
				StreamVideoColorRange        string   `json:"stream_video_color_range"`
				StreamVideoColorSpace        string   `json:"stream_video_color_space"`
				StreamVideoColorTrc          string   `json:"stream_video_color_trc"`
				StreamVideoDynamicRange      string   `json:"stream_video_dynamic_range"`
				StreamVideoHeight            string   `json:"stream_video_height"`
				StreamVideoWidth             string   `json:"stream_video_width"`
				StreamVideoRefFrames         string   `json:"stream_video_ref_frames"`
				StreamVideoLanguage          string   `json:"stream_video_language"`
				StreamVideoLanguageCode      string   `json:"stream_video_language_code"`
				StreamVideoScanType          string   `json:"stream_video_scan_type"`
				StreamVideoDecision          string   `json:"stream_video_decision"`
				StreamAudioBitrate           string   `json:"stream_audio_bitrate"`
				StreamAudioBitrateMode       string   `json:"stream_audio_bitrate_mode"`
				StreamAudioChannels          string   `json:"stream_audio_channels"`
				StreamAudioChannelLayout     string   `json:"stream_audio_channel_layout"`
				StreamAudioCodec             string   `json:"stream_audio_codec"`
				StreamAudioSampleRate        string   `json:"stream_audio_sample_rate"`
				StreamAudioChannelLayout0    string   `json:"stream_audio_channel_layout_"`
				StreamAudioLanguage          string   `json:"stream_audio_language"`
				StreamAudioLanguageCode      string   `json:"stream_audio_language_code"`
				StreamAudioDecision          string   `json:"stream_audio_decision"`
				StreamSubtitleCodec          string   `json:"stream_subtitle_codec"`
				StreamSubtitleContainer      string   `json:"stream_subtitle_container"`
				StreamSubtitleFormat         string   `json:"stream_subtitle_format"`
				StreamSubtitleForced         int      `json:"stream_subtitle_forced"`
				StreamSubtitleLocation       string   `json:"stream_subtitle_location"`
				StreamSubtitleLanguage       string   `json:"stream_subtitle_language"`
				StreamSubtitleLanguageCode   string   `json:"stream_subtitle_language_code"`
				StreamSubtitleDecision       string   `json:"stream_subtitle_decision"`
				StreamSubtitleTransient      int      `json:"stream_subtitle_transient"`
			} `json:"sessions"`
			StreamCountDirectPlay   int `json:"stream_count_direct_play"`
			StreamCountDirectStream int `json:"stream_count_direct_stream"`
			StreamCountTranscode    int `json:"stream_count_transcode"`
			TotalBandwidth          int `json:"total_bandwidth"`
			LanBandwidth            int `json:"lan_bandwidth"`
			WanBandwidth            int `json:"wan_bandwidth"`
		} `json:"data"`
	} `json:"response"`
}
