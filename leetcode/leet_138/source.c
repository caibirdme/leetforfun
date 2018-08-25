#include <stdio.h>
#include <stdlib.h>

struct RandomListNode
{
    int label;
    struct RandomListNode *next;
    struct RandomListNode *random;
};

struct RandomListNode *createNewNode(int label)
{
    struct RandomListNode *p = (struct RandomListNode *)malloc(sizeof(struct RandomListNode));
    p->label = label;
    p->random = NULL;
    p->next = NULL;
    return p;
}

struct RandomListNode *copyRandomList(struct RandomListNode *head)
{
    if (head == NULL)
    {
        return NULL;
    }
    struct RandomListNode *cur = head;
    while (cur != NULL)
    {
        struct RandomListNode *p = cur->next;
        cur->next = createNewNode(cur->label);
        cur->next->next = p;
        cur = p;
    }
    cur = head;
    while (cur != NULL)
    {
        if (cur->random != NULL)
        {
            cur->next->random = cur->random->next;
        }
        cur = cur->next->next;
    }
    struct RandomListNode *newHead = head->next;
    cur = head;
    while (cur != NULL)
    {
        struct RandomListNode *p = cur->next;
        cur->next = p->next;
        if (p->next != NULL)
        {
            p->next = p->next->next;
        }
        cur = cur->next;
    }
    return newHead;
}