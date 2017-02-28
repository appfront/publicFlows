#required params:
  #"message ="
  #"channel = "
  #"hook = "


  - task: slack
    image: appfront/post_to_slack
    vars:
    - CHANNEL=fwx_ops
    - MESSAGE=fwx_bo_st build successful
    - HOOK=https://hooks.slack.com/services/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
