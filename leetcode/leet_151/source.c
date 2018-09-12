#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void findWord(char *s, int start, int *realStart, int *length)
{
    while (s[start] != '\0' && s[start] == ' ')
    {
        start++;
    }
    *realStart = start;
    int count = 1;
    while (s[start + 1] != '\0' && s[start + 1] != ' ')
    {
        start++;
        count++;
    }
    *length = count;
    return;
}

void reverseStr(char *s, int start, int last)
{
    char t;
    while (start < last)
    {
        t = s[start];
        s[start] = s[last];
        s[last] = t;
        start++;
        last--;
    }
}

void moveStr(char *s, int start, int length, int step)
{
    if (step <= 0)
    {
        return;
    }
    int last = start + length - 1;
    while (start <= last)
    {
        s[start - step] = s[start];
        start++;
    }
}

void reverseWords(char *s)
{
    if (s == NULL)
    {
        return;
    }
    int last = strlen(s) - 1;
    if (last < 0)
    {
        return;
    }
    int i = 0;
    int j = last;
    while (i <= last && s[i] == ' ')
        i++;
    while (j > i && s[j] == ' ')
        j--;
    int n = i;
    if (n > 0 || j < last)
    {
        for (; i <= j; i++)
            s[i - n] = s[i];
    }

    if (j - n + 1 <= last)
    {
        s[j - n + 1] = '\0';
    }
    last = j - n;
    int count = 0;
    int start, length;
    reverseStr(s, 0, last);
    i = 0;
    while (i <= last)
    {
        //find consecutive space
        while (i <= last && s[i] == ' ')
        {
            i++;
            count++;
        }
        if (i > last)
        {
            s[i - count] = '\0';
            return;
        }
        findWord(s, i, &start, &length);
        i = start + length - 1;
        reverseStr(s, start, i);
        moveStr(s, start, length, count);
        int nextSpace = i - count + 1;
        if (nextSpace < last)
        {
            s[nextSpace] = ' ';
        }
        i += 2;
    }
    if (count > 0)
    {
        s[last - count + 1] = '\0';
    }
}

int main(int argc, char const *argv[])
{
    // int len = strlen(argv[1]);
    // char *data = (char *)malloc(len + 1);
    // strcpy(data, argv[1]);
    // char *s = data;
    // reverseWords(s);
    // printf("%s\n", s);
    char data[] = "   a    b ";
    char *s = data;
    reverseWords(s);
    printf("%s %lu\n", s, strlen(s));
    return 0;
}
