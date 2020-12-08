Attribute VB_Name = "NewMacros"
Sub AOC8_p1()
    Dim acc As Integer: acc = 0
    
    Dim file_name As String: file_name = "C:\Users\stdaha1\Downloads\input.txt"
    Dim strTextLine As String
    Dim my_file As Integer: my_file = FreeFile
    Open file_name For Input As #my_file

    Dim text_line As String
    Line Input #my_file, text_line
    
    Dim rows() As String
    rows = Split(text_line, vbLf)
    Dim length As Integer: length = UBound(rows) + 1
    
    Dim position As Integer: position = 0
    Dim i As Integer: i = 0
    For i = 0 To length
      If (StrComp("done", rows(position), vbTextCompare) = 0) Then
        MsgBox (acc)
        Exit For
      End If
            
      Dim instruction() As String: instruction = Split(rows(position), " ")
      Dim operation As String: operation = instruction(0)
      Dim value As String: value = CInt(instruction(1))
      
      rows(position) = "done"
      
      If (StrComp("nop", operation, vbTextCompare) = 0) Then
        position = position + 1
      ElseIf (StrComp("acc", operation, vbTextCompare) = 0) Then
        position = position + 1
        acc = acc + value
      ElseIf (StrComp("jmp", operation, vbTextCompare) = 0) Then
        position = position + value
      End If
      
    Next i
    
End Sub

Sub AOC8_p2()
    Dim acc As Integer: acc = 0
    Dim file_name As String: file_name = "C:\Users\stdaha1\Downloads\input.txt"
    Dim strTextLine As String
    Dim my_file As Integer: my_file = FreeFile
    Open file_name For Input As #my_file

    Dim text_line As String
    Line Input #my_file, text_line
    Dim rows() As String
    rows = Split(text_line, vbLf)
    Dim length As Integer: length = UBound(rows) + 1
    
    Close my_file
    
    Dim solved As Boolean: solved = False
    Dim position As Integer: position = 0
    Dim i As Integer: i = 0
    For i = 0 To length
    
      Dim instruction() As String: instruction = Split(rows(position), " ")
      Dim operation As String: operation = instruction(0)
      Dim value As String: value = CInt(instruction(1))
      
      If (StrComp("nop", operation, vbTextCompare) = 0) Then
        If Not solved Then
          If reach(position) = True Then
            solved = True
            MsgBox ("reached")
            position = position + value
          Else
            position = position + 1
          End If
        Else
          position = position + 1
        End If
        
      ElseIf (StrComp("acc", operation, vbTextCompare) = 0) Then
        position = position + 1
        acc = acc + value
        
      ElseIf (StrComp("jmp", operation, vbTextCompare) = 0) Then
        If Not solved Then
          If Not solved And reach(position) = True Then
            solved = True
            position = position + 1
          Else
            position = position + value
          End If
        Else
          position = position + value
        End If
      End If
      
      If (CInt(position) > (length - 1)) Then
        Exit For
      End If
      
    Next i
    MsgBox (acc)
End Sub

Function reach(switched_position As Integer) As Boolean
    If switch_position >= 0 Then
    
    Dim file_name As String: file_name = "C:\Users\stdaha1\Downloads\input.txt"
    Dim strTextLine As String
    Dim my_file As Integer: my_file = FreeFile
    Open file_name For Input As #my_file

    Dim text_line As String
    Line Input #my_file, text_line
    Dim rows() As String
    rows = Split(text_line, vbLf)
    Dim length As Integer: length = UBound(rows) + 1
    
    Close my_file
    
    Dim position As Integer: position = switched_position
    
    Dim instruction() As String: instruction = Split(rows(position), " ")
    Dim operation As String: operation = instruction(0)
    Dim value As String: value = instruction(1)
    If (StrComp("nop", operation, vbTextCompare) = 0) Then
      operation = "jmp"
    Else
      operation = "nop"
    End If
    rows(position) = operation & " " & value
    
    Dim solved As Boolean: solved = False
    Dim i As Integer: i = 0
    For i = 0 To length
      If (CInt(position) > length - 1) Then
        reach = True
        Exit For
      End If
      instruction = Split(rows(position), " ")
      
      operation = instruction(0)
      value = CInt(instruction(1))
      
      If (StrComp("nop", operation, vbTextCompare) = 0) Then
        position = position + 1
        
      ElseIf (StrComp("acc", operation, vbTextCompare) = 0) Then
        position = position + 1
        
      ElseIf (StrComp("jmp", operation, vbTextCompare) = 0) Then
        position = position + value
      End If
      
    Next i
    End If
    
End Function
